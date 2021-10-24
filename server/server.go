package main

import (
	"io"

	pb "github.com/asbeeq/grpc/pb/numbers"

	"fmt"
	"net"

	"google.golang.org/grpc"
)

type NumServer struct {
	pb.UnsafeNumServiceServer
}

func (n *NumServer) Sum(stream pb.NumService_SumServer) error {
	var total int64 = 0
	var counter int = 0
	for {
		next, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Received %d numbers sum: %d\n", counter, total)
			stream.SendAndClose(&pb.NumResponse{Total: total})
			return nil
		}
		if err != nil {
			return err
		}

		total = total + next.X
		counter++
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterNumServiceServer(s, &NumServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
