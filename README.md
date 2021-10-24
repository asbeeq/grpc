# grpc
go run main.go

---

to generate proto files:
- protoc --go_out=./user user.proto

---
to make protoc work:
- export GO_PATH=~/go
- export PATH=$PATH:/$GO_PATH/bin
