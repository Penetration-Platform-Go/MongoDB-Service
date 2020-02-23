package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Penetration-Platform-Go/MongoDB-Service/model"
	mongodb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
	"go.mongodb.org/mongo-driver/bson"
)

// QueryViews handle
func QueryViews(condition []*mongodb.Value) []model.Views {
	filter := make(map[string]interface{})
	for _, con := range condition {
		filter[con.Key] = con.Value
	}
	cursor, err := model.Query("Platform", "Views", filter)
	if err != nil {
		fmt.Println(err)
		return []model.Views{}
	}
	var results []model.Views
	for cursor.Next(context.TODO()) {
		var elem model.Views
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println(err)
		return []model.Views{}
	}
	cursor.Close(context.TODO())
	return results
}

// AddViews handle
func AddViews() {
	result := QueryViews([]*mongodb.Value{
		{Key: "key", Value: time.Now().Format("20000101")},
	})
	if len(result) == 0 {
		model.Insert("Platform", "Views", bson.M{
			"key":   time.Now().Format("20000101"),
			"value": 1,
		})
	} else {
		model.Update("Platform", "Views", bson.M{
			"$inc": bson.M{"value": 1},
		}, bson.M{
			"key": time.Now().Format("20000101"),
		})
	}
	model.Update("Platform", "Views", bson.M{
		"$inc": bson.M{"value": 1},
	}, bson.M{
		"key": "all",
	})
}
