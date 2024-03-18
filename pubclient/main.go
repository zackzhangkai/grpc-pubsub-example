package main

import (
	"context"
	"log"

	"github.com/zackzhangkai/grpc-pubsub-example/api/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewPubSubServiceClient(conn)

	_, err = client.Publish(context.Background(), &proto.PublishRequest{Topic: "gocn", Payload: []byte("hello, gophers!")})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &proto.PublishRequest{Topic: "greeting", Payload: []byte("hello, world!")})
	if err != nil {
		log.Fatal(err)
	}
}
