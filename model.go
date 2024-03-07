package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Auther_Id   string             `bson:"auther_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}
