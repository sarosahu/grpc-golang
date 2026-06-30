package main

import (
	"fmt"
	"log"

	pb "github.com/sarosahu/grpc-golang/greet/proto"
	"google.golang.org/grpc"
)

/*
 * Server streaming API
 */
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes() function was invoked : %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}