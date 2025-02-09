package main

import (
	"log"
	"time"

	pb "github.com/OmRamanuj/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error{
	log.Printf("Got Reuquested with Names : %v",req.Names)
	for _,Names:= range req.Names{
		res := &pb.HelloResponse{
			Message: "Hello "+Names,
			}
			if err := stream.Send(res); err != nil {
				return err
			}
			time.Sleep(2 * time.Second)
		}
		return nil
	}

