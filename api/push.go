package api

import (
	"delay_queue/grpc"
	"delay_queue/logic"
	"fmt"
	"golang.org/x/net/context"
)

func checkPush(req *delayqueue.PushRequest) (err error) {
	if req.Data == "" || req.DelaySeconds < 0 {
		return errParams
	}
	if len(req.NotifyUrl) > notifyUrlMaxLength {
		return errNotifyLength
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
	dataId, err := logic.Push(req.Data, req.DelaySeconds, req.NotifyUrl)
	if err != nil {
		return
	}
	reply.DataId = dataId
	return
}
