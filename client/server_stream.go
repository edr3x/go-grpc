package main

import (
	"context"
	"io"
	"log"

	pb "github.com/edr3x/go-grpc/proto"
)

func callSayhelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
		}

		log.Printf("Response from server: %v", message)
	}

	log.Printf("Streaming finished")
}
