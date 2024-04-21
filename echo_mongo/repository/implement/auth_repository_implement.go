package implement

import (
	"context"
	dto "echo_mongo/dto/auth"
	model "echo_mongo/model/user"
	"echo_mongo/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthImplement struct {
	mongoDB *mongo.Database
}

func NewAuthImplement(mongoDB *mongo.Database) repository.IAuthRepository {
	return &AuthImplement{
		mongoDB: mongoDB,
	}
}

func (i *AuthImplement) Login(ctx context.Context, dto dto.LoginDto) (*model.User, error) {
	collection := i.mongoDB.Collection("user")

	var user *model.User

	filter := bson.M{
		"username": dto.Username,
		"password": dto.Password,
	}

	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return &model.User{}, err
		}

		return &model.User{}, err
	}

	return user, nil
}

func (i *AuthImplement) Register(ctx context.Context, dto model.User) (*model.User, error) {
	collection := i.mongoDB.Collection("user")

	data := bson.M{
		"username":  dto.Username,
		"password":  dto.Password,
		"email":     dto.Email,
		"role":      "USER",
		"countryId": "1",
	}

	result, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	var user *model.User

	if err := collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}
