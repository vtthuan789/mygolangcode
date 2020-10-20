package main

import (
	"context"
	"fmt"
	"grpc_go/blog/blogpb"
	"io"
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
	blog1 := &blogpb.Blog{
		AuthorId: "Cap",
		Content:  "My name is Cap",
		Title:    "My second blog",
	}
	blog2 := &blogpb.Blog{
		AuthorId: "Hulk",
		Content:  "My name is Hulk",
		Title:    "My third blog",
	}
	blog3 := &blogpb.Blog{
		AuthorId: "Thor",
		Content:  "My name is Thor",
		Title:    "My fourth blog",
	}
	blogID := createBlog(c, blog)
	createBlog(c, blog1)
	createBlog(c, blog2)
	createBlog(c, blog3)
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
	listBlog(c)
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

func listBlog(bc blogpb.BlogServiceClient) {
	fmt.Println("Listing blogs from MongoDB")

	stream, err := bc.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while listing blogs from MongoDB: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while reading blog stream:", err)
		}
		fmt.Println(res.GetBlog())
	}
}
