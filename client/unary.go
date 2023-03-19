package main

import (
	"context"
	"log"
	"time"

	pb "github.com/edr3x/go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not call SayHello: %v", err)
	}

	log.Printf("Response from Unary call: %v", res)
}
