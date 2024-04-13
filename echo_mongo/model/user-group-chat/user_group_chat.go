package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserGroupChat struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserId      string             `json:"userId" bson:"userId"`
	GroupChatId string             `json:"groupChatId" bson:"groupChatId"`
}
