package main

import (
	"context"
	"log"
	"time"

	pb "github.com/OmRamanuj/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient,names *pb.NamesList){
	log.Printf("Client Streaming Started")
	stream,err:= client.SayHelloClientStreaming(context.Background())
	if err!=nil{
		log.Fatalf("Could not send names: %v", err)
	}
	for _,name:= range names.Names{
		req:= &pb.HelloRequest{
			Name:name,
		}
		if err:=stream.Send(req);err!=nil{
			log.Fatalf("Error while sending request : %v",err)
		}
		log.Printf("Sent name : %v",name)
		time.Sleep(2*time.Second)
}

	res,err := stream.CloseAndRecv()
	log.Printf("Client Streaming Ended")
	if err!=nil{
		log.Fatalf("Error while receiving response : %v",err)
	}
	log.Printf("%v",res.Messages)
}