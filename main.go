package main

import (
	"fmt"

	"github.com/asbeeq/grpc/user"

	"google.golang.org/protobuf/proto"
)

func main() {
	u := user.User{Name: "asdfg", UserId: "1"}
	fmt.Println("To encode:", u.String())
	encoded, err := proto.Marshal(&u)
	if err != nil {

		panic(err)
	}

	v := user.User{}
	err = proto.Unmarshal(encoded, &v)
	if err != nil {
		panic(err)
	}
	fmt.Println("Recovered:", v.String())
}
