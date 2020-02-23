package model

import (
	"context"

	mongodb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
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
	ID    string        `bson:"_id"`
	User  string        `bson:"user"`
	Title string        `bson:"title"`
	Score int32         `bson:"score"`
	IP    []*mongodb.Ip `bson:"ip"`
	Map   *mongodb.Map  `bson:"map"`
}

// Views define
type Views struct {
	Key   string `bson:"key"`
	Value int32  `bson:"value"`
}
