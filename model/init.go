package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient define
var MongoClient *mongo.Client

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, _ := mongo.Connect(context.TODO(), clientOptions)

	MongoClient = client
}

// Project define
type Project struct {
	ID    string `bson:"_id"`
	User  string `bson:"user"`
	Value string `bson:"value"`
}
