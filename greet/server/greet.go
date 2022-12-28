package main

import (
	"context"
	"log"
	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error){
	log.Printf("Greet Function was invoked with %v\n",in)
	return &pb.GreetResponse{
		Result:"Hello "+in.FirstName,
	},nil
}