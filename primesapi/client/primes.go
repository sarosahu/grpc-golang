package main

import (
	"context"
	"io"
	"log"

	pb "github.com/sarosahu/grpc-golang/primesapi/proto"
)

func doPrimes(c pb.PrimesAPIServiceClient) {
	log.Println("doPrimes() was invoked")

	stream, err := c.Primes(context.Background(), &pb.PrimesAPIRequest{
		Number: 120,
	})
	if err != nil {
		log.Fatalf("Error while calling PrimesRequest(): %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("Primes: %d\n", msg.Result)
	}
}