package main

import (
	"context"
	"log"

	pb "github.com/sarosahu/grpc-golang/greet/proto"
)

// We need to implement greet end point

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}