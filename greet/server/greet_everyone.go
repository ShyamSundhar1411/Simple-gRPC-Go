package main

import (
	"io"
	"log"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)

func (s *Server)GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error{
	log.Printf("Invoked GreetEveryone Function")
	for{
		request,err:=stream.Recv()
		if err==io.EOF{
			return nil
		}
		if err!=nil{
			log.Fatalf("Error reading : %v\n",err)
		}
		res:="Hello"+request.FirstName+"!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err!=nil{
			log.Fatalf("Error while sending data %v\n",err)
		}
	}
	return nil
}