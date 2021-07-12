package main

import (
	"context"
	"grpc-app/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	log.Printf("Add(%d, %d) = %d", x, y, x+y)
	result := x + y
	response := &proto.AddResponse{Result: result}
	return response, nil
}

func (server *server) Average(averageStream proto.AppService_AverageServer) error {
	total := int64(0)
	count := int64(0)
	for {
		avgReq, err := averageStream.Recv()
		if err == io.EOF {
			average := total / count
			log.Printf("Average(%d) = %d", count, average)
			averageResponse := &proto.AverageResponse{Result: average}
			averageStream.SendAndClose(averageResponse)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		total += avgReq.GetNo()
		count++
	}
	return nil
}
func main() {
	listener, err := net.Listen("tcp", "localhost:52000")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
