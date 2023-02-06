package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	hello "github.com/jsainthilaire/try-gripmock/proto/gen/go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:4770", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("dialing: %v", err)
	}
	defer conn.Close()

	client := hello.NewGreeterClient(conn)

	resp, err := client.Hi(context.Background(), &hello.HiRequest{Name: "Jose"})
	if err != nil {
		log.Fatalf("error from grpc: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetReply())
}
