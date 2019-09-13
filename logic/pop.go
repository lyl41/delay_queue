package logic

import (
	"delay_queue/common"
	"delay_queue/redis"
)

func Pop(timeout int64) (data string, err error) {
	payloadKey, err := redis.PopReadyQueue(common.QueueName, int(timeout))
	if err != nil {
		return
	}
	data, err = redis.GetPayload(payloadKey)
	if err != nil {
		return
	}
	_ = redis.DelPayload(payloadKey)
	return
}
