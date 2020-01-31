package grpc

import (
	"context"

	"github.com/Penetration-Platform-Go/MongoDB-Service/controller"
	"github.com/Penetration-Platform-Go/MongoDB-Service/model"
	mongodb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
)

// MongoDBService define
type MongoDBService struct {
}

// QueryProjectsByUsername method Query Projects By Username
func (u *MongoDBService) QueryProjectsByUsername(user *mongodb.Username, stream mongodb.MongoDB_QueryProjectsByUsernameServer) error {
	result := controller.QueryProject(user.Username)
	for _, val := range result {
		if err := stream.Send(&mongodb.ProjectInformation{
			Id:   val.ID,
			User: val.User,
			Value: &mongodb.ProjectValue{
				Temp: val.Value,
			},
		}); err != nil {
			return err
		}
	}

	return nil
}

// InsertProject method
func (u *MongoDBService) InsertProject(ctx context.Context, project *mongodb.ProjectInformation) (*mongodb.Result, error) {
	result := controller.InsertProject(&model.Project{
		User:  project.User,
		Value: project.Value.Temp,
	})
	return &mongodb.Result{
		IsVaild: result,
		Value:   "Failed",
	}, nil
}

// UpdateProject method
func (u *MongoDBService) UpdateProject(ctx context.Context, project *mongodb.ProjectInformation) (*mongodb.Result, error) {
	result := controller.UpdateProject(&model.Project{
		ID:    project.Id,
		User:  project.User,
		Value: project.Value.Temp,
	})
	return &mongodb.Result{
		IsVaild: result,
		Value:   "Failed",
	}, nil
}

// DeleteProject method
func (u *MongoDBService) DeleteProject(ctx context.Context, project *mongodb.ProjectId) (*mongodb.Result, error) {
	result := controller.DeleteProject(project.Id)
	return &mongodb.Result{
		IsVaild: result,
		Value:   "Failed",
	}, nil
}
