package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/asbeeq/grpc/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 3 {
		log.Fatal("not enough aruments ")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	op := flag.Arg(2)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAdderClient(conn)

	var res *api.AddResponse

	switch op {
	case "+":
		res, err = c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
		if err != nil {
			log.Fatal(err)
		}
	case "-":
		res, err = c.Sub(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println(res.GetResult())
}
