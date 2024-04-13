package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroupChat struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
