package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)
func doGreetEveryone(c pb.GreetServiceClient){
	stream,err:=c.GreetEveryone(context.Background())
	if err!=nil{
		log.Fatalf("Error while connecting%v\n",err)
	}
	reqs:=[]*pb.GreetRequest{
		{FirstName:"Test1"},
		{FirstName:"Test2"},
		{FirstName: "Test3"},
	}
	waitChannel :=make(chan struct{})
	go func(){
		for _,req := range reqs{
			log.Printf("Sending request %v\n",req)
			stream.Send(req)
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()
	go func(){
		for{
			res,err:=stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil{
				log.Printf("Error while receiving %v\n",err)
				break
			}
			log.Printf("Received: %v\n",res.Result)
		}
		close(waitChannel)
	}()
	<-waitChannel
}