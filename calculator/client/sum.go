package main

import (
	"context"
	"log"

	pb "github.com/sarosahu/grpc-golang/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked from client")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 8,
		SecondNumber: 13,
	})

	if err != nil {
		log.Fatal("Could not greet: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}