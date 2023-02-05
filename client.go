package main

import (
	"context"
	"log"

	"github.com/rohitv5/grpc_golang_basic/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client",
	}

	resp,err := c.SayHello(context.Background(),&message)

	log.Printf("Response from server: %v",resp.Body)
}
