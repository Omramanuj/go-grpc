package main

import (
	"io"
	"log"
	"context"
	pb "github.com/OmRamanuj/go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Streaming Started")
	stream,err:=client.SayHelloServerStreaming(context.Background(),names)
	if err!=nil{
		log.Fatalf("Could not send names: %v", err)
	}

	for{
		message,err := stream.Recv()
		if err == io.EOF{
			break
		}

		if err!=nil{
			log.Fatalf("Error while streaming : %v",err)
		}

		log.Println(message)
	}
	log.Printf("Streaming Ended")
}