package main

import (
	"flag"

	"github.com/rnsasg/go_grpc/chat/client"
	"github.com/rnsasg/go_grpc/chat/server"
)

var port = flag.Int("port", 50001, "gRPC Server Port")

func main() {

	// Start gRPC Server
	go server.Start(50001)

	// Start gRPC Client
	gcli := client.Start(50001)

	client.Greet(gcli, "Hello Server")
}
