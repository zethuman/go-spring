package main

import (
	"log"
	"net"
	"protobuf/pkg/api"
	"protobuf/pkg/reverser"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &reverser.GRPCServer{}

	api.RegisterReverserServer(s, srv)

	t, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(t); err != nil {
		log.Fatal(err)
	}
}
