package main

import (
	"log"

	pb "github.com/sarosahu/grpc-golang/primesapi/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(in *pb.PrimesAPIRequest, stream grpc.ServerStreamingServer[pb.PrimesAPIResponse]) error {
	log.Printf("Primes() was invoked by %v\n", in)

	num := in.Number
	var k int32 = 2
	for num > 1 {
		if num % k == 0 {
			//log.Printf("%d\n", k)
			num /= k
			stream.Send(&pb.PrimesAPIResponse{
				Result: k,
			})
		} else {
			k += 1
		}
	}
	return nil
}
