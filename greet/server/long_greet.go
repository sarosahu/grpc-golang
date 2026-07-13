package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/sarosahu/grpc-golang/greet/proto"
	"google.golang.org/grpc"
)

// Client streaming
func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("LongGreet() was invoked for server")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatal("Error while reading client stream: %v\n", err)
		}

		log.Printf("[Server]Receiving : %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}