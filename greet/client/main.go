package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:50051"

func main() {
	connection, err := grpc.Dial(address,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("Failed to connect: %v\n",err)
	}
	defer connection.Close()
}