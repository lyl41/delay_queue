package daemon

import (
	"delay_queue/redis"
	"fmt"
	"strconv"
	"time"
)

const (
	//zrange(0, DetectStop), http://doc.redisfans.com/sorted_set/zrange.html
	DetectStop = 5
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
	payloadKeys, err := redis.RangeZset(0, DetectStop)
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
			err = redis.PushReadyQueue(payloadKeys[i-1])
			if err != nil {
				fmt.Println("Detector PushReadyQueue err, ", err)
				continue
			}
			fmt.Println("success push to ready queue,", payloadKeys[i-1], val, time.Now().Unix())
			payloadKeysNeedDel = append(payloadKeysNeedDel, payloadKeys[i-1])
		}
	}
	if len(payloadKeysNeedDel) > 0 {
		if err = redis.RemZset(payloadKeysNeedDel); err != nil {
			fmt.Println("Detector RemZset err:", err)
		}
	}
}
