package logic

import (
	"delay_queue/common/util"
	"delay_queue/redis"
)

func Push(value string, TTR int64) (err error) {
	key := util.RandomStr(16) //generate payload key or id
	err = redis.SetPayload(key, value)
	if err != nil {
		return
	}
	err = redis.AddZset(key, TTR)
	if err != nil {
		return
	}
	return
}
