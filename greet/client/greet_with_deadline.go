package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func doGreetwithDeadline(c pb.GreetServiceClient,timeout time.Duration){
	log.Println("doGreet was invoked")
	ctx,cancel := context.WithTimeout(context.Background(),timeout)
	defer cancel()
	request := &pb.GreetRequest{
		FirstName: "Test1",
	}
	response,err:=c.GreetWithDeadline(ctx,request)
	if err!=nil{
		e,ok := status.FromError(err)
		if ok{
			if e.Code() == codes.DeadlineExceeded{
				log.Println("Deadline was exceeded !")
				return
			}else{
				log.Printf("Unexpected gRPC Error %v",err)
			}
		}else{
			log.Fatalf("A non gRPC Error %v",err)
		}
	}
	log.Printf("GreetwithDeadline: %v",response.Result)
}