package main
import (
	"context"
	"log"
	pb "github.com/ShyamSundhar1411/Simple-gRPC-Go/greet/proto"
)
func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	response,err:=c.Greet(context.Background(),&pb.GreetRequest{
		FirstName: "Shyam",
	})
	if err!=nil{
		log.Fatalf("Could not greet: %v\n",err)
	}
	log.Printf("Greetings: %s\n",response.Result)
}