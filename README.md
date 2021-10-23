# grpc
server - go run cmd/server/main.go

client - go run cmd/client/main.go 1 2 +

---

to generate proto files:
- protoc --go_out=./pkg/api --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./pkg/api api/proto/adder.proto
