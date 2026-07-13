package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sarosahu/grpc-golang/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage() was invoked from Client")

	reqs := []*pb.AvgApiRequest{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
	}
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatal("Error while calling Avg() from client %v\n", err)
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

	log.Printf("Avg: %f\n", res.Result)
}