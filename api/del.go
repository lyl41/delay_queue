package api

import (
	"delay_queue/grpc"
	"delay_queue/logic"
	"fmt"
	"golang.org/x/net/context"
)

func checkDelReq(req *delayqueue.DelRequest) (err error) {
	if req.DataId == "" {
		return errParams
	}
	return
}

func (Server) Del(ctx context.Context, req *delayqueue.DelRequest) (reply *delayqueue.DelReply, err error) {
	defer func() {
		if err != nil {
			fmt.Println("Pop err:", err)
		}
	}()
	reply = new(delayqueue.DelReply)
	if err = checkDelReq(req); err != nil {
		return
	}
	err = logic.DealDel(req.DataId)
	return
}
