package util

import (
	"math/rand"
	"time"
)

const (
	baseStr = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`
	baseNum = `1234567890`
)

var (
	limitStr       = int64(len(baseStr))
	limitNumberStr = int64(len(baseNum))
)

//产生随机字符串，包括英文大小写字母与数字
func RandomStr(length int) string {
	randomStr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		c := byte(baseStr[rand.Int63()%limitStr])
		randomStr = append(randomStr, c)
	}
	return string(randomStr)
}

//产生随机数字形式的字符串
func RandomNumberStr(length int) string {
	randomNumberStr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		c := byte(baseNum[rand.Int63()%limitNumberStr])
		randomNumberStr = append(randomNumberStr, c)
	}
	return string(randomNumberStr)
}
