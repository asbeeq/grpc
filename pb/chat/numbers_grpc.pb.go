// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SendTxt(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendTxtClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendTxt(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendTxtClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/numbers.ChatService/SendTxt", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendTxtClient{stream}
	return x, nil
}

type ChatService_SendTxtClient interface {
	Send(*ChatRequest) error
	Recv() (*StatsResponse, error)
	grpc.ClientStream
}

type chatServiceSendTxtClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendTxtClient) Send(m *ChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendTxtClient) Recv() (*StatsResponse, error) {
	m := new(StatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SendTxt(ChatService_SendTxtServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SendTxt(ChatService_SendTxtServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTxt not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SendTxt_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendTxt(&chatServiceSendTxtServer{stream})
}

type ChatService_SendTxtServer interface {
	Send(*StatsResponse) error
	Recv() (*ChatRequest, error)
	grpc.ServerStream
}

type chatServiceSendTxtServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendTxtServer) Send(m *StatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendTxtServer) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "numbers.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendTxt",
			Handler:       _ChatService_SendTxt_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "numbers.proto",
}
