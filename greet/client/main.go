package main

import (
	"log"
	"time"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)	

var address string = "localhost:50051"

func main() {
	tls := true
	options:=[]grpc.DialOption{}
	if tls{
		if tls{
			certFile := "ssl/ca.crt"

			creds,err := credentials.NewClientTLSFromFile(certFile,"")
	
			if err!=nil{
				log.Fatalf("Failed loading Ca Trust Certificates %v\n",err)
			}
			options = append(options,grpc.WithTransportCredentials(creds))
		}
	}
	connection, err := grpc.Dial(address,options...)
	if err!=nil{
		log.Fatalf("Failed to connect: %v\n",err)
	}
	defer connection.Close()
	client := pb.NewGreetServiceClient(connection)
	doGreetEveryone(client)
	doGreetwithDeadline(client,5*time.Second)
}