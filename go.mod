module github.com/asbeeq/grpc

go 1.16

// protoc --go_out=./pkg/api --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./pkg/api api/proto/adder.proto

require (
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.25.0 // indirect
)
