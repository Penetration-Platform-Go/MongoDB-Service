package grpc

import (
	"context"

	"github.com/Penetration-Platform-Go/MongoDB-Service/controller"
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

	return nil, nil
}

// UpdateProject method
func (u *MongoDBService) UpdateProject(ctx context.Context, project *mongodb.ProjectInformation) (*mongodb.Result, error) {

	return nil, nil
}

// DeleteProject method
func (u *MongoDBService) DeleteProject(ctx context.Context, project *mongodb.ProjectId) (*mongodb.Result, error) {

	return nil, nil
}
