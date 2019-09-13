package redis

import "github.com/garyburd/redigo/redis"

func PushReadyQueue(queueName string, payloadKey string) (err error) {
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("lpush", queueName, payloadKey)
	if err != nil {
		return
	}
	return
}

func BatchPushReadyQueue(queueName string, keys []string) (err error) {
	if len(keys) == 0 {
		return
	}
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("lpush", redis.Args{}.Add(queueName).AddFlat(keys)...)
	return
}

//timeout is seconds which command 'brpop' will block when queue is empty.
func PopReadyQueue(queueName string, timeout int) (payloadKey string, err error) {
	con := pool.Get()
	defer con.Close()
	nameAndData, err := redis.Strings(con.Do("brpop", queueName, timeout))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	if len(nameAndData) > 1 {
		payloadKey = nameAndData[1]
	}
	return
}
