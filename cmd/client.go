package cmd

import (
	"context"
	"io"
	"log"

	"grpcdemo/rpc"

	"google.golang.org/grpc"
)

func RunClient() {
	log.Println("starting client")
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// Close the connection once we're done.
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	client := rpc.NewDemoServiceClient(conn)

	request := &rpc.HelloWorldRequest{Name: "Warren", NickName: "seagray"}
	helloWorld(client, request)
	spellMyName(client, request)
}

func spellMyName(client rpc.DemoServiceClient, request *rpc.HelloWorldRequest) {
	stream, err := client.SpellMyName(context.Background(), request)
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

func helloWorld(client rpc.DemoServiceClient, request *rpc.HelloWorldRequest) {
	response, err := client.HelloWorld(context.Background(), request)
	if err != nil {
		panic(err)
	}

	log.Println(response.Greeting)
}
