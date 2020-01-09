package server

import (
	"context"
	"fmt"

	"grpcdemo/rpc"
)

type DemoService struct {
	rpc.UnimplementedDemoServiceServer
}

func (*DemoService) HelloWorld(_ context.Context, req *rpc.HelloWorldRequest) (*rpc.HelloWorldResponse, error) {
	return &rpc.HelloWorldResponse{Greeting: fmt.Sprintf("Hi, %s!", req.Name)}, nil
}
