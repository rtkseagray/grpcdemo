package cmd

import (
	"context"
	"io"
	"log"

	"grpcdemo/client"
	"grpcdemo/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunClient() {
	log.Println("starting client")

	// Note that our client wiring has changed slightly. We're configuring gRPC to send credentials
	// on each request from data stored in the request Context.
	conn, err := grpc.Dial(
		"localhost:7777",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(credentials.PerRPCCredentials(client.ContextBasedJWT{})),
	)

	if err != nil {
		panic(err)
	}

	// Close the connection once we're done.
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	c := rpc.NewDemoServiceClient(conn)

	request := &rpc.HelloWorldRequest{Name: "Warren", NickName: "seagray"}
	helloWorld(c, request)
	spellMyName(c, request)
}

func spellMyName(c rpc.DemoServiceClient, request *rpc.HelloWorldRequest) {
	stream, err := c.SpellMyName(client.Auth(context.Background(), "wgray"), request)
	if err != nil {
		panic(err)
	}

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		log.Println(recv.Letter)
	}

	if err = stream.CloseSend(); err != nil {
		panic(err)
	}
}

func helloWorld(c rpc.DemoServiceClient, request *rpc.HelloWorldRequest) {
	response, err := c.HelloWorld(client.Auth(context.Background(), "seagray"), request)
	if err != nil {
		panic(err)
	}

	log.Println(response.Greeting)
}
