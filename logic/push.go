package logic

import (
	"delay_queue/common"
	"delay_queue/common/util"
	"delay_queue/redis"
	"time"
)

func Push(value string, delaySeconds int64, notifyUrl string) (payloadKey string, err error) {
	key := util.RandomStr(common.PayloadKeyLength) //generate payload key or id
	if notifyUrl != "" {
		key = key + common.KeySep + notifyUrl //len(key) > 16
	}
	err = redis.SetPayload(key, value)
	if err != nil {
		return
	}
	ttr := time.Now().Unix() + delaySeconds
	err = redis.AddZset(key, ttr)
	if err != nil {
		return
	}
	payloadKey = key
	return
}
