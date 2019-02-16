package api

import (
	"delay_queue/grpc"
	"delay_queue/logic"
	"fmt"
	"golang.org/x/net/context"
)

func checkPop(req *delayqueue.PopRequest) (err error) {
	if req.Timeout < 0 {
		return errParams
	}
	return
}

func (Server) Pop(ctx context.Context, req *delayqueue.PopRequest) (reply *delayqueue.PopReply, err error) {
	defer func() {
		if err != nil {
			fmt.Println("Pop err:", err)
		}
	}()
	reply = new(delayqueue.PopReply)
	if err = checkPop(req); err != nil {
		return
	}
	reply.Data, err = logic.Pop(req.Timeout)
	if err != nil {
		return
	}
	fmt.Println("Pop return data", reply.Data)
	return
}
