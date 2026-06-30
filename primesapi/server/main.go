package main

import (
	"log"
	"net"

	pb "github.com/sarosahu/grpc-golang/primesapi/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50053"

type Server struct {
	pb.PrimesAPIServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterPrimesAPIServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}