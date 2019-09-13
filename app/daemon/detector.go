package daemon

import (
	"delay_queue/app/common"
	"delay_queue/app/redis"
	"fmt"
	"time"
)

func Detect() {
	fmt.Println("Detector running...")
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()
	for {
		<-ticker.C
		detect()
	}
}

func detect() {
	//取出所有比当前时间戳小的数据id
	payloadKeysAndScores, err := redis.RangeZsetByScore(0, time.Now().Unix())
	if err != nil {
		fmt.Println("Detector zrangebyscore err: ", err)
		return
	}
	if len(payloadKeysAndScores) == 0 {
		return
	}
	payloadKeysNeedDel := make([]string, 0, len(payloadKeysAndScores))
	notifyItems := make([]string, 0)
	items := make([]string, 0)
	for i := 0; i < len(payloadKeysAndScores); i += 2 { //奇数项是key
		key := payloadKeysAndScores[i]
		payloadKeysNeedDel = append(payloadKeysNeedDel, key)
		if len(payloadKeysAndScores[i]) == common.PayloadKeyLength {
			items = append(items, key)
		} else {
			notifyItems = append(notifyItems, key)
		}
	}
	err = redis.BatchPushReadyQueue(common.NotifyQueueName, notifyItems)
	if err != nil {
		fmt.Println("Detector BatchPushReadyQueue err, ", err)
		return
	}
	err = redis.BatchPushReadyQueue(common.QueueName, items)
	if err != nil {
		fmt.Println("Detector BatchPushReadyQueue err, ", err)
		return
	}
	fmt.Println("success push to ready queue, num:", len(payloadKeysNeedDel), time.Now().Unix())
	if err = redis.RemZset(payloadKeysNeedDel); err != nil {
		fmt.Println("Detector RemZset err:", err)
	}
}
