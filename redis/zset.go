package redis

import "github.com/garyburd/redigo/redis"

func AddZset(payloadKey string, score int64) (err error) {
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("zadd", ZsetName, score, payloadKey)
	if err != nil {
		return
	}
	return
}

func RemZset(payloadKeys []string) (err error) {
	if len(payloadKeys) == 0 {
		return
	}
	con := pool.Get()
	defer con.Close()
	_, err = con.Do("zrem", redis.Args{}.Add(ZsetName).AddFlat(payloadKeys)...)//TODO 这个点易错。
	return
}

//index sorted set from start to end, [start:end], eg: [0:1] will return[member1, score1, member2, score2]
func RangeZset(start, end int) (payloadKeys []string, err error) {
	con := pool.Get()
	defer con.Close()
	payloadKeys, err = redis.Strings(con.Do("zrange", ZsetName, start, end, "withscores"))
	return
}