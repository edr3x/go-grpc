package main

import (
	"context"
	"log"
	"time"

	pb "github.com/edr3x/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client streamging started")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Message: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Could not send request: %v", err)
		}

		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Printf("Client streaming finished")

	if err != nil {
		log.Fatalf("Could not receive response: %v", err)
	}

	log.Printf("Response: %v", res)
}
