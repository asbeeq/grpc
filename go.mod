module github.com/asbeeq/grpc

go 1.16

// client streaming example

// subdir "/numbers" created automatically, but "/pb" - manually
// problem with package underscore (package _) solved, after generation package name in pb files - "package numbers"
// protoc -I=. --go_out=./pb --go-grpc_out=./pb *.proto

// to make protoc work:
// export GO_PATH=~/go
// export PATH=$PATH:/$GO_PATH/bin

require (
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.25.0 // indirect
)
