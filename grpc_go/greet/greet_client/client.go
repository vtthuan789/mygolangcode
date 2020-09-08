package main

import (
	"fmt"
	"grpc_go/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Printf("Create client: %f", c)
}
