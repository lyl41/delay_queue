package redis

import (
	"fmt"
	"testing"
)

func TestPushReadyQueue(t *testing.T) {
	err := PushReadyQueue("queue_val")
	fmt.Println(err)
}

func TestPopReadyQueue(t *testing.T) {
	payloadKeys, err := PopReadyQueue(3)
	fmt.Println(payloadKeys)
	fmt.Println("err:", err)
}