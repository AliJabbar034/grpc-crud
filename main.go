package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	blogpb "githu.com/alijabbar034/proto"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	port = 8000
)

type Server struct {
	collection *mongo.Collection
	blogpb.UnimplementedBlogServiceServer
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := &Server{}
	url := os.Getenv("URL")
	if err := server.connectToDb(url); err != nil {
		fmt.Println("connection to db failed")
		panic(err)
		return
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println("error creating server..")
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	blogpb.RegisterBlogServiceServer(grpcServer, server)

	go func() {
		fmt.Println("Starting.................")
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println("error suring starting service..")
			panic(err)
		}
	}()

	cha := make(chan os.Signal, 1)

	signal.Notify(cha, os.Interrupt)
	<-cha

	fmt.Println("stopping server")
	grpcServer.Stop()
	lis.Close()

}

func (s *Server) connectToDb(url string) error {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	s.collection = client.Database("blogs").Collection("blog")

	fmt.Println("conection success")
	return nil
}
