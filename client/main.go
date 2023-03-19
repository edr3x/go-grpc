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

	defer conn.Close() // info: close the connection when the function returns

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"obi", "wan", "kenobi"},
	}

	// callSayHello(client) // info: unary call i.e. request, response

	// callSayhelloServerStream(client, names) //info: server streaming call i.e. request, stream of responses

	// callSayHelloClientStream(client, names) //info: client streaming call i.e. stream of requests, response

	callSayHelloBidirectionalStreaming(client, names) //info: bidirectional streaming call i.e. stream of requests, stream of responses
}
