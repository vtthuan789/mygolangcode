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
	blogID := createBlog(c, blog)
	readBlog(c, "5f7c941117f45dd1c348d0a9")
	readBlog(c, blogID)

	updatedBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Content:  "Changed Content",
		Title:    "Changed Title",
	}
	updateBlog(c, updatedBlog)
	readBlog(c, blogID)
	deleteBlog(c, blogID)
}

func createBlog(bc blogpb.BlogServiceClient, b *blogpb.Blog) string {
	fmt.Println("Creating a blog")

	blogReq := &blogpb.CreateBlogRequest{
		Blog: b,
	}

	res, err := bc.CreateBlog(context.Background(), blogReq)
	if err != nil {
		log.Fatalln("Error while creating blog:", err)
	}

	fmt.Println("Blog has been created:", res)
	return res.GetBlog().GetId()
}

func readBlog(bc blogpb.BlogServiceClient, blogID string) {
	fmt.Println("Reading a blog")

	res, err := bc.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: blogID,
	})
	if err != nil {
		fmt.Println("Error while reading blog:", err)
	} else {
		fmt.Println("Bog was read:", res)
	}
}

func updateBlog(bc blogpb.BlogServiceClient, blog *blogpb.Blog) {
	fmt.Println("Updating a blog:", blog)

	res, err := bc.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		fmt.Println("Error while updating blog:", err)
	} else {
		fmt.Println("Bog was updated:", res)
	}
}

func deleteBlog(bc blogpb.BlogServiceClient, blogID string) {
	fmt.Println("Deleting a blog")

	res, err := bc.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogID,
	})
	if err != nil {
		fmt.Println("Error while deleting blog:", err)
	} else {
		fmt.Println("Bog was deleted:", res)
	}
}
