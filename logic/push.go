package logic

import (
	"delay_queue/common"
	"delay_queue/common/util"
	"delay_queue/redis"
)

func Push(value string, TTR int64, notifyUrl string) (err error) {
	key := util.RandomStr(common.PayloadKeyLength) //generate payload key or id
	if notifyUrl != "" {
		key = key + common.KeySep + notifyUrl //len(key) > 16
	}
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
