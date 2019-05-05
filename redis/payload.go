package redis

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

func SetPayload(key, value string) (err error) {
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("set", key, value)
	if err != nil {
		return
	}
	return
}

func GetPayload(payloadKey string) (payload string, err error) {
	con := pool.Get()
	defer con.Close()
	payload, err = redis.String(con.Do("get", payloadKey))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	return
}

func DelPayload(payloadKey string) (err error) {
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("del", payloadKey)
	return
}

func SetMultiPayload(kv []*RedisKv) (successCount int, err error) { //TODO 优化一下批量查询，pipeline之类的。
	con := pool.Get()
	defer con.Close()
	for _, val := range kv {
		_, err = con.Do("set", val.Key, val.Value)
		if err != nil {
			continue
		}
		successCount++
	}
	return
}

const successPrefix = `success`

//const failPrefix = `fail`

func IncrRedisCount(userId int64, success bool) (err error) {
	prefix := successPrefix
	if !success {
		prefix = failPrefix
	}
	con := pool.Get()
	defer con.Close()
	key := prefix + strconv.FormatInt(userId, 10)
	count, err := redis.Int64(con.Do("incr", key))
	if err != nil {
		return
	}
	if count == 1 {
		timeout := 60
		if !success {
			timeout = 60
		}
		_, err = con.Do("expire", key, timeout)
		return
	}
	return
}
