package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc_go/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Sum() function is calling with", req)

	res := &calculatorpb.SumResponse{
		Sum: int64(req.GetOperand().GetOperand1()) + int64(req.GetOperand().GetOperand2()),
	}

	return res, nil
}

func main() {
	fmt.Println("Server is running!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
