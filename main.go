package main

import (
	"os"

	"grpcdemo/cmd"
)

//go:generate protoc demo.proto --go_out=plugins=grpc:rpc

func main() {
	if len(os.Args) == 2 {
		cmd.RunServer()
	} else {
		cmd.RunClient()
	}
}
