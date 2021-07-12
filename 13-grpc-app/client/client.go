package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:52000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewAppServiceClient(conn)
	doRequestResopnse(client)
}

func doRequestResopnse(client proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	response, err := client.Add(context.Background(), addRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.GetResult())
}
