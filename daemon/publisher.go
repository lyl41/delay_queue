package daemon

import (
	"delay_queue/common"
	"delay_queue/http_client"
	"delay_queue/redis"
	"fmt"
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
	url := payloadKey[common.PayloadKeyLength + 1:]
	err = http_client.SendPostRequest(url, payload) //TODO 添加重试
	if err != nil {
		fmt.Println("Publisher firstly send post err:", err)
		return
	}
}
