package api

import (
	"delay_queue/grpc"
	"delay_queue/logic"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func checkPush(req *delayqueue.PushRequest) (err error) {
	if req.Data == "" || req.Ttr < time.Now().Unix()-1 {
		return errParams
	}
	return
}

func (Server) Push(ctx context.Context, req *delayqueue.PushRequest) (reply *delayqueue.PushReply, err error) {
	defer func() {
		if err != nil {
			fmt.Println("Push err:", err)
		}
	}()
	reply = new(delayqueue.PushReply)
	if err = checkPush(req); err != nil {
		return
	}
	err = logic.Push(req.Data, req.Ttr, req.NotifyUrl)
	if err != nil {
		return
	}
	return
}
