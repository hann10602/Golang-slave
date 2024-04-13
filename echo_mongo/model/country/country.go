package country

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	code string             `json:"code" bson:"code"`
}
