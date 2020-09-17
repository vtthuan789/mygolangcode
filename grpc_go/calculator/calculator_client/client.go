package main

import (
	"context"
	"fmt"
	"grpc_go/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm a calculator client")

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect, error: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewSumServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.SumServiceClient) {
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
