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
	result := controller.QueryProjectByUsername(user.Username)
	for _, val := range result {
		if err := stream.Send(&mongodb.ProjectInformation{
			Id:    val.ID,
			User:  val.User,
			Score: val.Score,
			Ip:    val.IP,
			Map:   val.Map,
		}); err != nil {
			return err
		}
	}

	return nil
}

// QueryProjectsByID method Query Projects By Username
func (u *MongoDBService) QueryProjectsByID(ctx context.Context, projectId *mongodb.ProjectId) (*mongodb.ProjectInformation, error) {
	result := controller.QueryProjectByID(projectId.Id)
	return &mongodb.ProjectInformation{
		Id:    result.ID,
		User:  result.User,
		Score: result.Score,
		Ip:    result.IP,
		Map:   result.Map,
	}, nil

}

// InsertProject method
func (u *MongoDBService) InsertProject(ctx context.Context, project *mongodb.ProjectInformation) (*mongodb.Result, error) {
	result := controller.InsertProject(&model.Project{
		User:  project.User,
		Score: project.Score,
		IP:    project.Ip,
		Map:   project.Map,
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
		Score: project.Score,
		IP:    project.Ip,
		Map:   project.Map,
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
