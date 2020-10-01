package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"grpc_go/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet() function is calling with %v\n", req)
	result := "Hello " + req.GetGreeting().GetFirstName()
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes() function is calling with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet() function is calling with a streaming request")
	result := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// Finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalln("Error while reading client stream", err)
		}

		result += "Hi " + req.GetGreeting().GetFirstName() + "! "
	}
}

func (*server) GreetEveryone(
	stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone() function is calling with a streaming request")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln("Error while reading client stream:", err)
			return err
		}

		result := "Hello " + req.GetGreeting().GetFirstName() + "!"
		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})

		if err != nil {
			log.Fatalln("Error while sending response to client:", err)
			return err
		}
	}
}

func (*server) GreetWithDeadline(
	ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (
	*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline() function is calling with %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("Client canceled the request")
			return nil, status.Error(codes.Canceled,
				"The client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	result := "Hello " + req.GetGreeting().GetFirstName()
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalln("Error while loading certificates: ", sslErr)
		return
	}

	s := grpc.NewServer(grpc.Creds(creds))
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
