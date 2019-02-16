package redis

import "github.com/garyburd/redigo/redis"

func PushReadyQueue(payloadKey string) (err error) {
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("lpush", QueueName, payloadKey)
	if err != nil {
		return
	}
	return
}

//timeout is seconds which command 'brpop' will block when queue is empty.
func PopReadyQueue(timeout int) (payloadKey string, err error) {
	con := pool.Get()
	defer con.Close()
	nameAndData, err := redis.Strings(con.Do("brpop", QueueName, timeout))
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
