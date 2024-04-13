package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title,unique"`
	Description string             `json:"description" bson:"description,unique"`
	Content     string             `json:"content" bson:"content,unique"`
	UserId      string             `json:"userId" bson:"userId"`
}
