package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"grpc_go/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (
	*calculatorpb.SumResponse, error) {
	fmt.Println("Sum() function is calling with", req)

	res := &calculatorpb.SumResponse{
		Sum: int64(req.GetOperand().GetOperand1()) +
			int64(req.GetOperand().GetOperand2()),
	}

	return res, nil
}

func (*server) PrimeNumbersDecomposition(
	req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.CalculatorService_PrimeNumbersDecompositionServer) error {
	fmt.Println("PrimeNumbersDecomposition() function is calling with", req)

	var p int32 = 2
	c := req.GetCompositeNumber().GetValue()

	for c > 1 {
		if c%p == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				PrimeNumber: p,
			}
			stream.Send(res)
			time.Sleep(time.Second)
			c = c / p
		} else {
			p = p + 1
		}
	}

	return nil
}

func main() {
	fmt.Println("Server is running!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
