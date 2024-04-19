package implement

import (
	"context"
	dto "echo_mongo/dto/user"
	model "echo_mongo/model/user"
	"echo_mongo/repository"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserImplement struct {
	mongoDB *mongo.Database
}

func NewUserImplement(mongoDB *mongo.Database) repository.IUserRepository {
	return &UserImplement{
		mongoDB: mongoDB,
	}
}

func (u *UserImplement) GetUsers(ctx context.Context) ([]*model.User, error) {
	collection := u.mongoDB.Collection("user")

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var users []*model.User

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (u *UserImplement) GetUserById(ctx context.Context, id string) (*model.User, error) {
	collection := u.mongoDB.Collection("user")

	var user *model.User

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	if err := collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserImplement) CreateUser(ctx context.Context, dto model.User) (*model.User, error) {
	collection := u.mongoDB.Collection("user")

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

	user, err := u.GetUserById(ctx, fmt.Sprint(result.InsertedID))

	if err != nil {
		return nil, err
	}

	return user, nil

	// var user *model.User

	// if err := collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
	// 	return nil, err
	// }

	// return user, nil
}

func (u *UserImplement) UpdateUser(ctx context.Context, dto dto.UpdateUserDto, id string) (*model.User, error) {
	collection := u.mongoDB.Collection("user")

	data := bson.M{
		"$set": bson.M{
			"email": dto.Email,
			"role":  dto.Role,
		},
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectId}, data)

	if err != nil {
		return nil, err
	}

	var user *model.User

	if err := collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserImplement) DeleteUser(ctx context.Context, id string) error {
	collection := u.mongoDB.Collection("user")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		return err
	}

	return nil
}
