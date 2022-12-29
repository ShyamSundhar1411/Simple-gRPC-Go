package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)

func doLongMany(c pb.GreetServiceClient){
	log.Printf("GreetMany");
	reqs:=[]*pb.GreetRequest{
		{FirstName:"Test1"},
		{FirstName:"Test2"},
		{FirstName: "Test3"},
	}
	stream,err:=c.LongGreet(context.Background())
	if err!=nil{
		log.Fatalf("Error while fetching : %v\n",err)
	}
	for _,req:=range reqs{
		log.Printf("Sending the req: %v\n",req)
		stream.Send(req)
		time.Sleep(1*time.Second)
	}
	response,err:=stream.CloseAndRecv()
	if err!=nil{
		log.Fatalf("Error %v\n",err)
	}else{
		log.Println(response.Result)
	}
}