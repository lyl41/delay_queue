package redis

import (
	"fmt"
	"testing"
)

func TestSetPayload(t *testing.T) {
	if err := SetPayload("qwe", "value");err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}

//func TestSetMultiPayload(t *testing.T) {
//	args :=  []string{"lyl1", "lyl1val", "lyl2", "lyl2val", "lyl3", "lyl3val"}
//	if err := SetMultiPayload(args); err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println("ok")
//}