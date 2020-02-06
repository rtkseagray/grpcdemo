package cmd

import (
	"log"
	"net"

	"grpcdemo/rpc"
	"grpcdemo/server"

	"google.golang.org/grpc"
)

func RunServer() {
	log.Println("starting server")

	lis, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// We've wired up our middleware here. Though only one unary and one streaming interceptor can
	// be wired up to the server, they can be chained in their implementation (so one interceptor
	// can explicitly call the next).
	s := grpc.NewServer(grpc.UnaryInterceptor(server.UnaryLoggingMiddleware), grpc.StreamInterceptor(server.StreamingLoggingMiddleware))
	rpc.RegisterDemoServiceServer(s, &server.DemoService{})
	log.Fatal(s.Serve(lis))
}
