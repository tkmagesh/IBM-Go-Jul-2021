package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:52000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewAppServiceClient(conn)
	//doRequestResopnse(client)
	//doClientStreaming(client)
	doServerStreaming(client)
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

func doClientStreaming(client proto.AppServiceClient) {
	data := []int64{3, 5, 2, 1, 4}
	averageClientStream, err := client.Average(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range data {
		fmt.Println("Sending ", value)
		averageRequest := &proto.AverageRequest{No: value}
		averageClientStream.Send(averageRequest)
		time.Sleep(500 * time.Millisecond)
	}
	averageResponse, responseErr := averageClientStream.CloseAndRecv()
	if responseErr != nil {
		log.Fatalln(responseErr)
	}
	fmt.Println(averageResponse.GetResult())
}

func doServerStreaming(client proto.AppServiceClient) {
	req := &proto.PrimeNumberRequest{RangeStart: 20, RangeEnd: 100}
	resStream, err := client.GeneratePrime(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Prime # = ", res.GetPrimeNumber())
	}
}
