package main

import (
	"context"
	"fmt"

	blogpb "githu.com/alijabbar034/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting client .....................")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:8000", opts)
	if err != nil {
		fmt.Println("error during ...")
		panic(err)
	}

	c := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AutherId:    "234",
		Title:       "fisrt blog",
		Description: "blog description.",
	}

	createBlogRes, err := c.CreateBlog(context.TODO(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		fmt.Println("error during  creating blog", err.Error())
		return
	}
	find, err := c.FindBlog(context.TODO(), &blogpb.FindBlogRequest{Id: "65ea327ebcd17d73df99e885"})
	if err != nil {
		fmt.Println("error during find blog", err.Error())
		return
	}
	findAll, err := c.FindAllBLog(context.Background(), &blogpb.FindAllBlogRequest{})
	if err != nil {
		fmt.Println("error during  find all blog", err.Error())
		return
	}
	deleteOne, err := c.DeleteBlog(context.TODO(), &blogpb.DeleteBlogRequest{Id: "65ea327ebcd17d73df99e885"})
	if err != nil {
		fmt.Println("error during deleting blog", err.Error())
		return
	}

	if err != nil {
		fmt.Println("error during blog", err.Error())
		return
	}
	fmt.Printf("created successfully / %v", createBlogRes)
	fmt.Printf("find one \n %v", find)
	fmt.Printf("find All \n %v", findAll)
	fmt.Printf("deleted one  \n %v", deleteOne)

}
