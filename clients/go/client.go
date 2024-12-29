package main

import (
	"context"
	"log"

	"github.com/daheige/athena/clients/go/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8081"
	// address     = "localhost:50050" // nginx grpc_pass port
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterServiceClient(conn)

	// Contact the server and print out its response.
	res, err := c.SayHello(context.Background(), &pb.HelloReq{
		Id: 1,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("name:%s,message:%s", res.Name, res.Message)

	res2, err := c.SayHello(context.Background(), &pb.HelloReq{
		Id: 0,
	})

	log.Println(res2, err)

	res3, err := c.Info(context.Background(), &pb.InfoReq{
		Name: "daheige",
	})

	log.Println(res3, err)
}

/*
2024/12/23 22:11:37 name:hello,world,message:call ok
2024/12/23 22:11:37 <nil> rpc error: code = InvalidArgument desc = Key: 'HelloReq.Id'
Error:Field validation for 'Id' failed on the 'required' tag
*/
