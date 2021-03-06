package main

import (
	"context"
	"fmt"
	"grpc_go/calculator/calculatorpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm a calculator client")

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect, error: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	//doUnary(c)
	//doStreamingServer(c)
	//doStreamingClient(c)
	//doBiDiStreaming(c)
	doErrorHandling(c, -1)
	doErrorHandling(c, 1024)
	doErrorHandling(c, 1234)
}

func doErrorHandling(cc calculatorpb.CalculatorServiceClient, n int32) {
	fmt.Println("Start doing a Square Root RPC...")

	res, err := cc.SquareRoot(context.Background(),
		&calculatorpb.SquareRootRequest{
			Number: n,
		})

	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			fmt.Println(s.Message())
			fmt.Println(s.Code())
			if s.Code() == codes.InvalidArgument {
				fmt.Println("Request was sent with negative number!")
			}
		} else {
			log.Fatalln("Error from server: ", err)
		}
	} else {
		fmt.Println("Receive response from SquareRoot RPC: ", res)
	}
}

func doBiDiStreaming(cc calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start doing a BiDi Streaming Client RPC...")

	// Create stream by invoking client
	stream, err := cc.FindMaximum(context.Background())
	if err != nil {
		log.Fatalln("Error while calling FindMaximum RPC:", err)
	}

	// Send a bunch of requests
	go func() {
		for _, number := range []int32{-12, 34, -78, 56, 67, -89} {
			fmt.Println("Sending request:", number)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	// Receive a bunch of response
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while reading server response:", err)
			break
		}
		fmt.Println("Receive response from server:", res)
	}
}

func doStreamingClient(cc calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start doing a Average Streaming Client RPC...")

	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			Number: 5,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 6,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 7,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 8,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 898,
		},
	}

	stream, err := cc.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalln("Error while calling ComputeAverage RPC", err)
	}

	for _, req := range requests {
		fmt.Println("Sending request:", req)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while receiving server response from ComputeAverage", err)
	}

	fmt.Println("Recieve response from ComputeAverage RPC:", res.GetAverage())
}

func doStreamingServer(cc calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start doing a Streaming Prime Number Decomposition RPC...")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		CompositeNumber: &calculatorpb.CompositeNumber{
			Value: 31415936,
		},
	}

	res, err := cc.PrimeNumbersDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling PrimeNumbersDecomposition RPC: %v", err)
	}

	for {
		num, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error while reading stream:", err)
		}

		fmt.Println("Receive response from PrimeNumbersDecomposition service:",
			num.GetPrimeNumber())
	}
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start doing a Unary Sum RPC...")

	req := &calculatorpb.SumRequest{
		Operand: &calculatorpb.SumOperand{
			Operand1: 12,
			Operand2: 56,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	fmt.Println("Receive response from Calculator service:", res.Sum)
}
