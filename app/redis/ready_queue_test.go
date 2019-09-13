package redis

import (
	"delay_queue/app/common"
	"fmt"
	"testing"
)

func TestPushReadyQueue(t *testing.T) {
	err := PushReadyQueue(common.QueueName, "queue_val")
	fmt.Println(err)
}

func TestPopReadyQueue(t *testing.T) {
	payloadKeys, err := PopReadyQueue(common.QueueName, 3)
	fmt.Println(payloadKeys)
	fmt.Println("err:", err)
}

func TestBatchPushReadyQueue(t *testing.T) {
	err := BatchPushReadyQueue(common.NotifyQueueName, []string{"nn", "mm"})
	fmt.Println(err)
}
