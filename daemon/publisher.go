package daemon

import (
	"delay_queue/common"
	"delay_queue/http_client"
	"delay_queue/redis"
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
		if payloadKey != "" {
			go publish(payloadKey) //TODO use go routine pool
		}
	}
}

var PostFrequency = []int{0, 2, 8, 30, 60*2, 60*5, 60*30, 60*60}

func publish(payloadKey string) {
	if len(payloadKey) <= common.PayloadKeyLength {
		fmt.Println("Publisher payloadKey not valid, ", payloadKey)
		return
	}
	payload, err := redis.GetPayload(payloadKey)
	if err != nil {
		fmt.Println("Publisher getPayload err:", err)
		return
	}
	if len(payload) == 0 {
		return
	}
	url := payloadKey[common.PayloadKeyLength+1:]
	err = http_client.SendPostRequest(url, payload) //TODO 添加重试、删除记录
	if err != nil {
		count, e := handlePostErr(payloadKey)
		fmt.Println("Publisher send post count:", count, "err:", err, time.Now())
		if e != nil {
			fmt.Println("rePost fail,", e)
		}
		return
	}
	fmt.Println("post success", payloadKey)
}

func handlePostErr(payloadKey string) (count int, err error){
	//取失败计数
	count, err = redis.GetFailCount(payloadKey)
	if err != nil {
		return //暂时先不通知了，防止延时队列累积过多消息
	}
	count++
	if count >= len(PostFrequency) { //这里表明超出最大通知次数。
		return
	}
	//计算TTR，写入到zset中
	delay := PostFrequency[count]
	nextPostTime := time.Now().Unix() + int64(delay)
	//更新失败计数
	err = redis.SetFailCount(payloadKey, count, delay<<1)
	if err != nil {
		return
	}
	err = redis.AddZset(payloadKey, nextPostTime)
	if err != nil {
		return
	}
	return
}