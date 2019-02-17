package daemon

import (
	"delay_queue/common"
	"delay_queue/redis"
	"fmt"
	"strconv"
	"time"
)

func Detect() {
	fmt.Println("Detector running...")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		detect()
	}
}

func detect() {
	payloadKeys, err := redis.RangeZset(0, common.DetectStop)
	if err != nil {
		fmt.Println("Detector zrange err: ", err)
		return
	}
	timestamp := time.Now().Unix()
	payloadKeysNeedDel := make([]string, 0)
	for i, val := range payloadKeys {
		if (i & 1) == 0 { // even is member, odd is score
			continue
		}
		score, _ := strconv.ParseInt(val, 10, 64)
		if score <= timestamp { //Need to return, Push to ready queue.
			payloadKey := payloadKeys[i-1]
			queueName := common.NotifyQueueName
			if len(payloadKey) == common.PayloadKeyLength {
				queueName = common.QueueName
			}
			err = redis.PushReadyQueue(queueName, payloadKey)
			if err != nil {
				fmt.Println("Detector PushReadyQueue err, ", err)
				continue
			}
			fmt.Println("success push to ready queue,", payloadKey, val, time.Now().Unix())
			payloadKeysNeedDel = append(payloadKeysNeedDel, payloadKey)
		}
	}
	if len(payloadKeysNeedDel) > 0 {
		if err = redis.RemZset(payloadKeysNeedDel); err != nil {
			fmt.Println("Detector RemZset err:", err)
		}
	}
}
