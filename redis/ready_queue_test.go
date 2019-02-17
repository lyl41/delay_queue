package redis

import (
	"delay_queue/common"
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
