package api

import (
	"delay_queue"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Server struct {
}

const (
	notifyUrlMaxLength = 66
)

var (
	errParams       = errors.New("参数错误")
	errNotifyLength = errors.Errorf("notify_url最长不能超过%d", notifyUrlMaxLength)
)

func (Server) Ping(ctx context.Context, req *delayqueue.PingRequest) (*delayqueue.PingReply, error) {
	fmt.Println("recv Ping msg:" + req.Msg)
	return &delayqueue.PingReply{Msg: "PONG~"}, nil
}
