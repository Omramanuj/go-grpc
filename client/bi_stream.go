package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/OmRamanuj/go-grpc/proto"
)

func callSayHelloBiDiStream(client pb.GreetServiceClient,names *pb.NamesList){
	log.Printf("BiDi Streaming Started")
	stream,err:= client.SayHelloBiDiStreaming(context.Background())
	if err!=nil{
		log.Fatalf("Could not send names: %v", err)
	}
	waitc:=make (chan struct{})
	go func(){
		for{
			message,err:= stream.Recv()
			if err == io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("Error while streaming : %v",err)
			}
			log.Println(message)
		}
		close(waitc)
	
	}()

	for _,name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err :=stream.Send(req); err!=nil{
			log.Fatalf("Error while sending request : %v",err)
		}
		time.Sleep(2*time.Second)

	}
	stream.CloseSend()
	<-waitc
	log.Printf("BiDi Streaming Ended")
}
