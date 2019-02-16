package redis

import (
	"fmt"
	"testing"
)

func TestSetZset(t *testing.T) {
	err := AddZset("zetmember1", 1233)
	err = AddZset("zetmember2", 123)
	err = AddZset("zetmember3", 12)
	fmt.Println(err)
}

func TestDelZset(t *testing.T) {
	var zset = []string{"zetmember1"}
	err := RemZset(zset)
	fmt.Println(err)
}

func TestRangeZset(t *testing.T) {
	payloadKeys, err := RangeZset(0, 1)
	fmt.Println(payloadKeys, err)
}
