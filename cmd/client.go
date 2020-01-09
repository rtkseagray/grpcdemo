package cmd

import (
	"context"
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

	response, err := client.HelloWorld(context.Background(), &rpc.HelloWorldRequest{Name: "Warren", NickName: "seagray"})
	if err != nil {
		panic(err)
	}

	log.Println(response.Greeting)
}
