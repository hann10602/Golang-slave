package db

import (
	"context"
	"echo_mongo/initializer"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	DBName string
}

func (m *MongoDB) Connect() {
	config, err := initializer.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	clientOptions := options.Client().ApplyURI(config.UriAddress)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("MongoDB connect successfully")
	m.Client = client
}

func (m *MongoDB) Close() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("MongoDB connection closed")
}
