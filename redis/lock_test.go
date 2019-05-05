package redis

import (
	"fmt"
	"testing"
)

func TestGetRedisLock(t *testing.T) {
	lock, err := Lock("lyl_lock")
	fmt.Println(lock, err)
}

func TestSetRedisUnlock(t *testing.T) {
	err := Unlock("lyl_lock")
	fmt.Println(err)
}
