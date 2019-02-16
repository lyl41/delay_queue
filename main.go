package main

import (
	"delay_queue/api"
	"delay_queue/daemon"
	"delay_queue/grpc"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

func main() {
	address := ":9002"
	l, err := net.Listen("tcp", address) //TODO config_file
	if err != nil {
		panic(errors.New("监听" + address + "失败"))
	}
	go daemon.Detect()// TODO need graceful close？
	s := api.Server{}
	grpcServer := grpc.NewServer()
	delayqueue.RegisterDelayQueueServer(grpcServer, s)
	fmt.Println("grpc Server listening ",address + "...")
	if err = grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
