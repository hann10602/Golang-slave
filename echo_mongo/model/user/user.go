package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username,unique"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email,unique"`
	Role      string             `json:"role" bson:"role"`
	CountryId string             `json:"countryId" bson:"countryId"`
}
