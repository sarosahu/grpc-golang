package main

import (
	"io"
	"log"

	pb "github.com/sarosahu/grpc-golang/calculator/proto"
	grpc "google.golang.org/grpc"
)

func (s *Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	var maxNum int32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal("Error while reading client stream: %v\n", err)
		}
		res := max(maxNum, req.Num)

		if res > maxNum {
			err = stream.Send(&pb.MaxResponse{
				Result: res,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client : %v\n", err)
			}
			maxNum = res
		}
		
	}
	
}