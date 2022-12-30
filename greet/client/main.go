package main

import (
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)	

var address string = "localhost:50051"

func main() {
	connection, err := grpc.Dial(address,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("Failed to connect: %v\n",err)
	}
	defer connection.Close()
	client := pb.NewGreetServiceClient(connection)
	doGreetEveryone(client)
	doGreetwithDeadline(client,5*time.Second)
}