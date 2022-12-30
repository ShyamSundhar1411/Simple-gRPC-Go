package main

import (
	"log"
	"net"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)
var address string = "0.0.0.0:50051"
type Server struct{
	pb.GreetServiceServer
}
func main(){

	options:=[]grpc.ServerOption{}
	tls:=true
	if tls{
		certFile := "ssl/server.crt"
		key := "ssl/server.pem"
		creds,err := credentials.NewServerTLSFromFile(certFile,key)

		if err!=nil{
			log.Fatalf("Failed loading Certificates %v\n",err)
		}
		options = append(options,grpc.Creds(creds))
	}
	listener,err := net.Listen("tcp",address)
	if err!=nil{
		log.Fatalf("Failed to listen on: %v\n",err)
	}
	log.Printf("Listening on %s\n",address)

	serverInstance := grpc.NewServer(options...)
	pb.RegisterGreetServiceServer(serverInstance,&Server{})
	if err = serverInstance.Serve(listener); err !=nil{
		log.Fatalf("Failed to serve: %v\n",err)
	}
}