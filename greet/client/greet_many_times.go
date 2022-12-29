package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)

func doGreetMany(c pb.GreetServiceClient){
	log.Printf("GreetMany");
	request:=&pb.GreetRequest{
		FirstName: "Shyam",
	}
	stream,err :=c.GreetManyTimes(context.Background(),request)
	if err!=nil{
		log.Fatalf("Error while calling GreetMany: %v\n",err)
	}else{
		for {
			message,err:=stream.Recv()
			if err == io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("Error while reading: %v\n",err)
			}else{
				log.Printf("GreeManyTimes: %s\n",message)
			}
		}
	}
}