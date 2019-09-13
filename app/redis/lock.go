package redis

import (
	"github.com/garyburd/redigo/redis"
)

const (
	redisLockTimeout = 10 // 10 seconds
)

func Lock(key string) (isLock bool, err error) {
	con := pool.Get()
	defer con.Close()
	_, err = redis.String(con.Do("set", key, 1, "ex", redisLockTimeout, "nx"))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	isLock = true
	return
}

func Unlock(key string) (err error) {
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("del", key)
	if err != nil {
		return
	}
	return
}
