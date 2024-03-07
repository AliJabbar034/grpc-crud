package main

import (
	"context"
	"fmt"

	blogpb "githu.com/alijabbar034/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {

	blog := req.GetBlog()

	data := BlogItem{
		Auther_Id:   blog.AutherId,
		Title:       blog.Title,
		Description: blog.Description,
	}

	inserted, err := s.collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println("Error creating blog")
		return nil, err
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:          inserted.InsertedID.(primitive.ObjectID).Hex(),
			AutherId:    blog.AutherId,
			Title:       blog.Title,
			Description: blog.Description,
		},
	}, nil
}

func (s *Server) FindBlog(ctx context.Context, req *blogpb.FindBlogRequest) (*blogpb.FindBlogResponse, error) {

	id := req.GetId()
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	blogItem := &BlogItem{}
	filter := bson.M{"_id": _id}

	if err := s.collection.FindOne(context.Background(), filter).Decode(&blogItem); err != nil {
		return nil, err
	}

	return &blogpb.FindBlogResponse{
		Blog: &blogpb.Blog{
			Id:          blogItem.ID.Hex(),
			AutherId:    blogItem.Auther_Id,
			Title:       blogItem.Title,
			Description: blogItem.Description,
		},
	}, nil

}

func (s *Server) DeleteBlog(ctx context.Context, rq *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	id := rq.GetId()
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	deletedCount, err := s.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	return &blogpb.DeleteBlogResponse{
		Count:  deletedCount.DeletedCount,
		Messag: "deleted successfully",
	}, nil

}

func (s *Server) FindAllBLog(ctx context.Context, req *blogpb.FindAllBlogRequest) (*blogpb.FindAllBlogResponse, error) {

	var blogs []BlogItem

	filter := bson.M{}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}

	var response []*blogpb.Blog

	for _, blog := range blogs {
		res := &blogpb.Blog{
			Id:          blog.ID.Hex(),
			AutherId:    blog.Auther_Id,
			Title:       blog.Title,
			Description: blog.Description,
		}
		response = append(response, res)
	}
	return &blogpb.FindAllBlogResponse{
		Blog: response,
	}, nil

}
