package main

import (
	"log"
	"net"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
	"google.golang.org/grpc"
)
var address string = "0.0.0.0:50051"
type Server struct{
	pb.GreetServiceServer
}
func main(){
	listener,err := net.Listen("tcp",address)
	if err!=nil{
		log.Fatalf("Failed to listen on: %v\n",err)
	}
	log.Printf("Listening on %s\n",address)

	serverInstance := grpc.NewServer()
	pb.RegisterGreetServiceServer(serverInstance,&Server{})
	if err = serverInstance.Serve(listener); err !=nil{
		log.Fatalf("Failed to serve: %v\n",err)
	}
}