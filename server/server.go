package main

import (
	"fmt"
	"io"
	"time"

	pb "github.com/asbeeq/grpc/pb/chat"

	"net"

	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnsafeChatServiceServer
}

func (c *ChatServer) SendTxt(stream pb.ChatService_SendTxtServer) error {
	var total int64 = 0
	go func() {
		for {
			t := time.NewTicker(time.Second * 2)
			select {
			case <-t.C:
				stream.Send(&pb.StatsResponse{TotalChar: total})
			}
		}
	}()

	for {
		next, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Client closed")
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println("->", next.Txt)
		total = total + int64(len(next.Txt))
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &ChatServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
