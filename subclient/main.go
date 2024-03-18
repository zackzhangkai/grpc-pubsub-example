package main

import (
	"context"
	"github.com/zackzhangkai/grpc-pubsub-example/api/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewPubSubServiceClient(conn)
	subscribeClient, err := client.Subscribe(context.Background(), &proto.SubscribeRequest{Topic: "gocn"})
	if err != nil {
		return
	}

	for {
		resp, err := subscribeClient.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Println(string(resp.Payload))
	}
	defer conn.Close()
}
