package logic

import "delay_queue/redis"

func Pop(timeout int64) (data string, err error) {
	payloadKey, err := redis.PopReadyQueue(int(timeout))
	if err != nil {
		return
	}
	data, err = redis.GetPayload(payloadKey)
	if err != nil {
		return
	}
	return
}
