package main

import (
	"context"
	"fmt"
	"grpc_go/greet/greetpb"
	"io"
	"log"
	"time"

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
	doClientStreaming(c)
	doBiDiStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing a Client Streaming RPC...")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tony",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Harry",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sonic",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rio",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Suri",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalln("Error while calling LongGreet RPC", err)
	}

	// Send each message request individually
	for _, req := range requests {
		fmt.Println("Sending request:", req)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	// Close request stream and receive response
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while receiving server response from LongGreet:", err)
	}

	// Print result
	fmt.Println("Received response from LongGreet service:", res)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing a Bi-Directional Streaming RPC...")
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tony",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Harry",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sonic",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rio",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Suri",
			},
		},
	}

	// Create stream by invoking client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln("Error while calling GreetEveryone RPC", err)
	}

	waitc := make(chan struct{})
	// Send a bunch of messages to server
	go func() {
		// Send each message request individually
		for _, req := range requests {
			fmt.Println("Sending request:", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()
	// Recevie a bunch of messages from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(
					"Error while receiving server response from GreetEveryone:",
					err)
				break
			}
			fmt.Println("Received response from GreetEveryone service:", res)
		}
		close(waitc)
	}()
	// Block until everything is done
	<-waitc
}
