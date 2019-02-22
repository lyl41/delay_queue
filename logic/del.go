package logic

import "delay_queue/redis"

func DealDel(dataId string) (err error) {
	err = redis.DelPayload(dataId)
	if err != nil {
		return
	}
	var val = []string{dataId}
	err = redis.RemZset(val)
	return
}
