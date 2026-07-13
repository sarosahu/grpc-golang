package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/sarosahu/grpc-golang/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone() invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Saroj"},
		{FirstName: "Sachin"},
		{FirstName: "Rahul"},
	}

	waitc := make(chan struct{})

	// This go routine will send stream of requests
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep((1 * time.Second))
		}
		stream.CloseSend()
	}()

	// This go routine will receive stream of responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving : %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}