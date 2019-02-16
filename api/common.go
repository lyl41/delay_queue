package api

import (
	"delay_queue/grpc"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Server struct {
}

var (
	errParams = errors.New("参数错误")
)



func (Server) Del(context.Context, *delayqueue.DelRequest) (*delayqueue.DelReply, error) {
	panic("implement me")
}

func (Server) Ping(ctx context.Context, req *delayqueue.PingRequest) (*delayqueue.PingReply, error) {
	fmt.Println("recv Ping msg:" + req.Msg)
	return &delayqueue.PingReply{Msg:"PONG~"}, nil
}