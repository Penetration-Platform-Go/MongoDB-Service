package controller

import (
	"context"
	"log"

	"fmt"
	"github.com/Penetration-Platform-Go/MongoDB-Service/model"
	mongodb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// QueryProjects handle
func QueryProjects(condition []*mongodb.Value) []model.Project {
	filter := make(map[string]interface{})
	for _, con := range condition {
		filter[con.Key] = con.Value
		if con.Key == "_id" {
			ido, err := primitive.ObjectIDFromHex(con.Value)
			if err != nil {
				fmt.Println(err)
				return []model.Project{}
			}
			filter[con.Key] = ido
		}
	}

	cursor, err := model.Query("Platform", "Projects", filter)
	if err != nil {
		fmt.Println(err)
		return []model.Project{}
	}
	var results []model.Project
	for cursor.Next(context.TODO()) {
		var elem model.Project
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println(err)
		return []model.Project{}
	}
	cursor.Close(context.TODO())
	return results

}

// InsertProject handle
func InsertProject(project *model.Project) bool {
	err := model.Insert("Platform", "Projects", bson.M{
		"user":      project.User,
		"score":     0,
		"title":     project.Title,
		"equipment": project.Equipment,
		"map":       project.Map,
	})
	if err != nil {
		return false
	}
	return true
}

// UpdateProject handle
func UpdateProject(condition []*mongodb.Value, key []string, value bson.M) bool {
	filter := make(map[string]interface{})
	for _, con := range condition {
		filter[con.Key] = con.Value
		if con.Key == "_id" {
			ido, err := primitive.ObjectIDFromHex(con.Value)
			if err != nil {
				fmt.Println(err)
				return false
			}
			filter[con.Key] = ido
		}
	}
	val := make(map[string]interface{})
	for _, con := range key {
		val[con] = value[con]
	}
	return model.Update("Platform", "Projects", bson.M{
		"$set": val,
	}, filter)
}

// DeleteProject handle
func DeleteProject(condition []*mongodb.Value) bool {
	filter := make(map[string]interface{})
	for _, con := range condition {
		filter[con.Key] = con.Value
		if con.Key == "_id" {
			ido, err := primitive.ObjectIDFromHex(con.Value)
			if err != nil {
				fmt.Println(err)
				return false
			}
			filter[con.Key] = ido
		}
	}

	return model.Delete("Platform", "Projects", filter)

}
