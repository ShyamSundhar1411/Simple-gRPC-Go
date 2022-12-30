package main

import (
	"context"
	"log"
	"time"
	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Server)GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error){
	log.Printf("Greet with Deadline was invoked")
	for i:=0;i<2;i++{
		if ctx.Err() == context.DeadlineExceeded{
			log.Println("Client Cancelled the Request")
			return nil,status.Error(codes.Canceled,"The Client canclled the request")
		}
		time.Sleep(1*time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello "+in.FirstName,
	},nil
}