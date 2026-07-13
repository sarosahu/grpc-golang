package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sarosahu/grpc-golang/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet() was invoked from Client")

	reqs := []*pb.GreetRequest{
		{FirstName: "Saroj"},
		{FirstName: "Sachin"},
		{FirstName: "Rahul"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatal("Error while calling LongGreet() from client %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("[Client]Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}