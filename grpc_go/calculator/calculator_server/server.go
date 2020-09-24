package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) ComputeAverage(
	stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage() function is calling with a streaming request")
	var sum int32 = 0
	noOfRequests := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			average := float32(sum) / float32(noOfRequests)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: average,
			})
		}

		if err != nil {
			log.Fatalln("Error while reading client stream", err)
		}

		sum += req.GetNumber()
		noOfRequests++
	}

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
