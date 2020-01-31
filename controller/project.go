package controller

import (
	"context"
	"log"

	"github.com/Penetration-Platform-Go/MongoDB-Service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"fmt"
)

// QueryProject handle
func QueryProject(username string) []model.Project {
	cursor, err := model.Query("Platform", "Projects", bson.M{"user": username})
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
	err := model.Insert("Platform", "Projects", bson.D{
		{"user", project.User},
		{"value", project.Value},
	})
	if err != nil {
		return false
	}
	return true
}

// UpdateProject handle
func UpdateProject(project *model.Project) bool {
	id, err := primitive.ObjectIDFromHex(project.ID)
	if err != nil {
		return false
	}
	return model.Update("Platform", "Projects", bson.M{
		"$set": bson.M{
			"user":  project.User,
			"value": project.Value,
		},
	}, bson.M{"_id": id})
}

// DeleteProject handle
func DeleteProject(id string) bool {
	ido, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return model.Delete("Platform", "Projects", bson.M{"_id": ido})

}
