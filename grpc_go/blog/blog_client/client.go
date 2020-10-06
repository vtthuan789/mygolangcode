package main

import (
	"context"
	"fmt"
	"grpc_go/blog/blogpb"
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

	c := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AuthorId: "Tony",
		Content:  "My name is Tony",
		Title:    "My first blog",
	}
	createBlog(c, blog)
}

func createBlog(bc blogpb.BlogServiceClient, b *blogpb.Blog) {
	fmt.Println("Creating a blog")

	blogReq := &blogpb.CreateBlogRequest{
		Blog: b,
	}

	res, err := bc.CreateBlog(context.Background(), blogReq)
	if err != nil {
		log.Fatalln("Error while creating blog: ", err)
	}

	fmt.Println("Blog has been created: ", res)
}
