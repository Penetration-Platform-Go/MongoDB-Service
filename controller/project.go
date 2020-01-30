package controller

import (
	"context"
	"log"

	"github.com/Penetration-Platform-Go/MongoDB-Service/model"
	"go.mongodb.org/mongo-driver/bson"

	"fmt"
)

// QueryProject handle
func QueryProject(username string) []model.Project {
	cursor, err := model.Query("Platform", "Projects", &bson.D{{"user", username}})
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
