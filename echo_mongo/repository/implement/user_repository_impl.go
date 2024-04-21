package implement

import (
	"context"
	dto "echo_mongo/dto/user"
	model "echo_mongo/model/user"
	"echo_mongo/repository"

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

func (i *UserImplement) GetUsers(ctx context.Context) ([]*model.User, error) {
	collection := i.mongoDB.Collection("user")

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

func (i *UserImplement) GetUserById(ctx context.Context, id string) (*model.User, error) {
	collection := i.mongoDB.Collection("user")

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

func (i *UserImplement) UpdateUser(ctx context.Context, dto dto.UpdateUserDto, id string) (*model.User, error) {
	collection := i.mongoDB.Collection("user")

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
