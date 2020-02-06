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

// QueryProjects method Query Projects By Condition
func (u *MongoDBService) QueryProjects(condition *mongodb.Condition, stream mongodb.MongoDB_QueryProjectsServer) error {
	result := controller.QueryProjects(condition.Value)
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
func (u *MongoDBService) DeleteProject(ctx context.Context, condition *mongodb.Condition) (*mongodb.Result, error) {
	result := controller.DeleteProject(condition.Value)
	return &mongodb.Result{
		IsVaild: result,
		Value:   "Failed",
	}, nil
}
