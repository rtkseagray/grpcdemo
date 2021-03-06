package server

import (
	"context"
	"fmt"
	"time"

	"grpcdemo/rpc"
)

type DemoService struct {
	rpc.UnimplementedDemoServiceServer
	sleepTime time.Duration
}

// Note that the signature of our handler didn't change when we regenerated the server code.
func (*DemoService) HelloWorld(_ context.Context, req *rpc.HelloWorldRequest) (*rpc.HelloWorldResponse, error) {
	return &rpc.HelloWorldResponse{Greeting: fmt.Sprintf("Hi, %s! Should I call you %s?", req.Name, req.NickName)}, nil
}

func (d *DemoService) SpellMyName(req *rpc.HelloWorldRequest, srv rpc.DemoService_SpellMyNameServer) error {
	for _, c := range req.Name {
		time.Sleep(d.sleepTime)
		if err := srv.Send(&rpc.LetterResponse{Letter: string(c)}); err != nil {
			return err
		}
	}

	return nil
}
