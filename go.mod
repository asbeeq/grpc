module github.com/asbeeq/grpc

go 1.16

// protoc --go_out=./user user.proto

// to make protoc work:
// export GO_PATH=~/go
// export PATH=$PATH:/$GO_PATH/bin

require (
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.25.0 // indirect
)
