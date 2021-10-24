# grpc
server - go run cmd/server/main.go

client - go run cmd/client/main.go 1 2 +

---

to generate proto files:
- protoc --go_out=./pkg/api --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./pkg/api api/proto/adder.proto

---

in case this error happens: "Please specify a program using absolute path or make sure the program is available in your PATH system variable"

to make protoc work:
- export GO_PATH=~/go
- export PATH=$PATH:/$GO_PATH/bin
