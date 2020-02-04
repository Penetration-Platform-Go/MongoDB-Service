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
func Query(database string, collection string, filter bson.M) (*mongo.Cursor, error) {

	fmt.Print("Query: " + database + " " + collection + " with filter:")
	fmt.Println(filter)

	result, err := getCollection(database, collection).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Insert Method by database, collection and value
func Insert(database string, collection string, value bson.M) error {

	fmt.Print("Insert: " + database + " " + collection + " with value:")
	fmt.Println(value)

	_, err := getCollection(database, collection).InsertOne(context.TODO(), value)
	return err
}

// Update Method by database, collection, filter value
func Update(database string, collection string, value bson.M, filter bson.M) bool {

	fmt.Print("Update: " + database + " " + collection + " with filter:")
	fmt.Println(filter)
	fmt.Print("Value:")
	fmt.Println(value)

	result, err := getCollection(database, collection).UpdateOne(context.TODO(), filter, value)

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result.MatchedCount != 1 {
		return false
	}

	return true

}

// Delete Method by database, collection and filter
func Delete(database string, collection string, filter bson.M) bool {

	fmt.Print("Delete: " + database + " " + collection + " with filter:")
	fmt.Println(filter)

	result, err := getCollection(database, collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if result.DeletedCount != 1 {
		return false
	}

	return true

}
