package redis

import (
	"fmt"
	"testing"
)

func TestSetPayload(t *testing.T) {
	if err := SetPayload("qw e_https:// falj/232", "value"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}

func TestGetPayload(t *testing.T) {
	payload, err := GetPayload("qw e_https:// falj/232")
	fmt.Println(payload, err)
}

func TestDelPayload(t *testing.T) {
	fmt.Println(DelPayload("lyl"))
}

//func TestSetMultiPayload(t *testing.T) {
//	args :=  []string{"lyl1", "lyl1val", "lyl2", "lyl2val", "lyl3", "lyl3val"}
//	if err := SetMultiPayload(args); err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println("ok")
//}

func TestIncrRedisCount(t *testing.T) {
	err := IncrRedisCount(123, false)
	fmt.Println(err)
}
