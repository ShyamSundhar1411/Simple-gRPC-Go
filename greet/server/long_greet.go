package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error{
	log.Printf("Long Greet Function was invoked with %v\n",stream)
	response := ""
	for{
		request,err := stream.Recv()
		if err==io.EOF{
			return stream.SendAndClose(&pb.GreetResponse{
				Result: response,
			})
		}
		if err!=nil{
			log.Fatalf("Error while processing request: %v\n",err)
		}
		response+=fmt.Sprintf("Hello %s:\n",request.FirstName)
	}
	return nil
}