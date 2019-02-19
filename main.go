package main

import (
	"delay_queue/api"
	"delay_queue/daemon"
	"delay_queue/grpc"
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	address := ":9002"
	l, err := net.Listen("tcp", address) //TODO config_file
	if err != nil {
		panic(errors.New("监听" + address + "失败"))
	}
	go daemon.Detect() // TODO need graceful close？
	go daemon.Publish()

	//-----------------------------zipkin
	collector, err := zipkin.NewHTTPCollector("http://localhost:9411/api/v1/spans")
	if err != nil {
		log.Fatal(err)
		return
	}
	tracer, err := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, "localhost:0", "grpc_server"),
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	opentracing.InitGlobalTracer(tracer)
	//---------------------------------------------

	s := api.Server{}
	//--
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	delayqueue.RegisterDelayQueueServer(grpcServer, s)
	fmt.Println("grpc Server listening ", address+"...")
	if err = grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
