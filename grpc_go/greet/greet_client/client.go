package main

import (
	"context"
	"fmt"
	"grpc_go/greet/greetpb"
	"io"
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

	// fmt.Printf("Create client: %f", c)

	doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing a Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Tony",
			LastName:  "Stark",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalln("Error while calling Greet RPC:", err)
	}

	fmt.Println("Receive response from Greet service:", res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing a Sever Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Iron",
			LastName:  "Man",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalln("Error while calling GreetManyTimes RPC:", err)
	}

	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error while reading stream:", err)
		}

		fmt.Println("Receive response from GreetManyTimes service:", msg.GetResult())
	}

}
