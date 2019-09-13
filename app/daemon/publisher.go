package daemon

import (
	"delay_queue/app/common"
	"delay_queue/app/http_client"
	"delay_queue/app/redis"
	"fmt"
	"time"
)

func Publish() {
	fmt.Println("Publisher running...")
	for {
		//TODO check redis是否会在空闲时释放连接
		payloadKey, err := redis.PopReadyQueue(common.NotifyQueueName, common.PublisherPopQueueTimeout)
		if err != nil {
			fmt.Println("Publisher pop ready queue err:", err)
			continue
		}
		payload, err := redis.GetPayload(payloadKey)
		if err != nil {
			fmt.Println("Publisher getPayload err:", err)
			//出现了错误，我们重新扔回队列，保证消息不丢失
			err = redis.PushReadyQueue(common.NotifyQueueName, payloadKey)
			if err != nil {
				fmt.Println("getPayload err and Push ready queue err", err)
			}
			continue
		}
		//readyQueue中有重复值时，这里保证不会发送多次，弱保证。当被主动pop了时，也不会发送。
		if len(payload) == 0 {
			continue
		}
		//publish中如果有错误，会扔回zset
		go publish(payloadKey, payload)
	}
}

var PostFrequency = []int{0, 2, 8, 30, 60 * 2, 60 * 5, 60 * 30, 60 * 60}

func publish(payloadKey, payload string) {
	url := payloadKey[common.PayloadKeyLength+1:]
	err := http_client.SendPostRequest(url, payload) //TODO 创建全局的http client
	if err != nil {
		count, e := handlePostErr(payloadKey)
		fmt.Println("Publisher send post count:", count, "err:", err, time.Now())
		if e != nil {
			fmt.Println("rePost fail,", e)
		}
		return
	}
	fmt.Println("post success", payloadKey, time.Now())
	//处理成功才删除消息
	_ = redis.DelPayload(payloadKey)
}

func handlePostErr(payloadKey string) (count int, err error) {
	//取失败计数
	count, _ = redis.GetFailCount(payloadKey)
	count++
	if count >= len(PostFrequency) { //这里表明超出最大通知次数, 丢弃这条消息
		return
	}
	//计算TTR，写入到zset中
	delay := PostFrequency[count]
	nextPostTime := time.Now().Unix() + int64(delay)
	//更新失败计数
	err = redis.SetFailCount(payloadKey, count, delay<<1)
	if err != nil {
		//这里不丢弃消息
	}
	err = redis.AddZset(payloadKey, nextPostTime)
	if err != nil {
		return
	}
	return
}
