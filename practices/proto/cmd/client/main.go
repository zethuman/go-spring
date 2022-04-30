package main

import (
	"context"
	"flag"
	"log"
	"protobuf/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("please enter string as a flag argument")
	}

	text := flag.Arg(0)

	conn, err := grpc.Dial(":8010", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewReverserClient(conn)
	response, err := client.Reverse(context.Background(), &api.ReverseRequest{Reversable: text})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.GetReversed())

}