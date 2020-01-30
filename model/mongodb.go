package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// getCollection
func getCollection(database string, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}

// Query Method by database and collection
func Query(database string, collection string, filter *bson.D) (*mongo.Cursor, error) {

	fmt.Print("Query: " + database + " " + collection + " with filter:")
	fmt.Println(*filter)

	result, err := getCollection(database, collection).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
