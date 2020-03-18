package main

import (
	"context"
	"delay_queue"
	"delay_queue/app/api"
	"delay_queue/app/daemon"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	address := ":9002"
	l, err := net.Listen("tcp", address) //TODO config_file
	if err != nil {
		panic(errors.New("监听" + address + "失败"))
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go daemon.Detect(ctx, wg)
	wg.Add(1)
	go daemon.Publish(ctx, wg)
	wg.Add(1)
	go func() {
		s := api.Server{}
		grpcServer := grpc.NewServer()
		delayqueue.RegisterDelayQueueServer(grpcServer, s)
		fmt.Println("grpc Server listening ", address+"...")
		go func() {
			time.Sleep(time.Second) // insure started
			<-ctx.Done()
			fmt.Println("grpc server try to stop...")
			grpcServer.GracefulStop()
			fmt.Println("grpc server stopped")
			wg.Done()
		}()
		if err = grpcServer.Serve(l); err != nil {
			panic(err)
		}
	}()
	si := <-sig
	fmt.Println("server recv stop signal", si)
	cancel()
	fmt.Println("server  waiting all deamon stopping...")
	wg.Wait()
	fmt.Println("daemon all stopped, now server stopped. bye~")
}
