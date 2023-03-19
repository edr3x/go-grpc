package main

import (
	pb "github.com/edr3x/go-grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// callSayHello(client) // info: unary call i.e. request, response

	names := &pb.NamesList{
		Names: []string{"obi", "wan", "kenobi"},
	}
	callSayhelloServerStream(client, names) //info: server streaming call i.e. request, stream of responses
}
