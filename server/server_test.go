package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"grpcdemo/rpc"
)

type ServerTestSuite struct {
	suite.Suite
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, &ServerTestSuite{})
}

// Testing unary methods are fairly straightforward. Simply construct a request object, call your
// handler endpoint, and verify the response.
func (s *ServerTestSuite) TestHelloWorld() {
	server := &DemoService{}
	request := rpc.HelloWorldRequest{Name: "Warren", NickName: "Seagray"}
	response, err := server.HelloWorld(context.Background(), &request)

	s.Assert().NoError(err)
	s.Assert().Equal(
		fmt.Sprintf("Hi, %s! Should I call you %s?", request.Name, request.NickName),
		response.Greeting,
	)
}

// Streaming RPCs are more complicated to test. This is an example of using gRPC's buffcon package.
// buffcon implements an in-memory full-duplex network connection using channels. We can start the
// server on a buffcon listener and then connect the generated client code to it to verify the
// server handlers.
func (s *ServerTestSuite) TestSpellMyName() {

	// Create the server and bind it to the buffcon.
	listener := bufconn.Listen(3)
	server := grpc.NewServer()
	service := DemoService{sleepTime: 10 * time.Millisecond}

	rpc.RegisterDemoServiceServer(server, &service)
	go func() {
		err := server.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
	defer server.Stop()

	// Connect our generated gRPC client to the server.
	conn, err := grpc.Dial(
		"",
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (conn net.Conn, err error) {
				return listener.Dial()
			},
		),
		grpc.WithInsecure(),
	)

	if err != nil {
		panic(err)
	}
	client := rpc.NewDemoServiceClient(conn)

	// Create a request and send it to the server.
	request := rpc.HelloWorldRequest{Name: "AB"}
	stream, err := client.SpellMyName(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	for _, l := range request.Name {
		recv, err := stream.Recv()
		s.Assert().NoError(err)
		s.Assert().Equal(string(l), recv.Letter)
	}

	_, err = stream.Recv()
	s.Assert().Equal(io.EOF, err)
}

func (s *ServerTestSuite) SetupSuite() {
	panic("implement me")
}
