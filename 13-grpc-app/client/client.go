package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:52000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewAppServiceClient(conn)
	//doRequestResopnse(client)
	//doClientStreaming(client)
	//doServerStreaming(client)
	//doBidirectionalStreaming(client)
	doRequestResopnseWithTimeout(client)
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

func doBidirectionalStreaming(client proto.AppServiceClient) {
	//bidirectional streaming
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	requests := []*proto.GreetEveryoneRequest{
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Magesh",
				LastName:  "Kuppan",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Ramesh",
				LastName:  "Jayaraman",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Suresh",
				LastName:  "Kannan",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Rajesh",
				LastName:  "Kumar",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "John",
				LastName:  "Philip",
			},
		},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message %v\n", req.Greeting)
			stream.Send(req)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("Received message %v\n", res.GetMessage())
		}
		close(waitc)
	}()
	<-waitc
}

func doRequestResopnseWithTimeout(client proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := client.Add(ctx, addRequest)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline exceeded.")
			} else {
				log.Fatalln(statusErr)
			}
		} else {
			log.Fatalf("error while calling Greet : %v\n", err)
		}
		return
	}
	fmt.Println(response.GetResult())
}
