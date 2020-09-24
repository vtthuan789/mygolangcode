package main

import (
	"context"
	"fmt"
	"grpc_go/calculator/calculatorpb"
	"io"
	"log"
	"time"

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

	doUnary(c)
	doStreamingServer(c)
	doStreamingClient(c)
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
