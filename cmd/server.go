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

	s := grpc.NewServer()
	rpc.RegisterDemoServiceServer(s, &server.DemoService{})
	log.Fatal(s.Serve(lis))
}
