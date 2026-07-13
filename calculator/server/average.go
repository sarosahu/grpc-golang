package main

import (
	"io"
	"log"

	pb "github.com/sarosahu/grpc-golang/calculator/proto"
	grpc "google.golang.org/grpc"
)

func (s *Server) Avg(stream grpc.ClientStreamingServer[pb.AvgApiRequest, pb.AvgApiResponse]) error {
	log.Println("Avg() is invoked to Server")

	var sum int32 = 0
	//var res float32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			//avg := float64(sum)/float64(count)
			return stream.SendAndClose(&pb.AvgApiResponse{
				Result: float64(sum)/float64(count),
			})
		}

		if err != nil {
			log.Fatal("Error while reading client stream: %v\n", err)
		}

		log.Printf("[Server]Receiving : %d\n", req)
		//res += fmt.Sprintf("Hello %s!\n", req.Num)
		sum += req.Num
		count++
	}
}