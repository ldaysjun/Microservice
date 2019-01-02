package handlers

import (
	"context"

	pb "trss"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.TimServiceServer {
	return timserviceService{}
}

type timserviceService struct{}

// TimedTask implements Service.
func (s timserviceService) TimedTask(ctx context.Context, in *pb.TimedTaskRequest) (*pb.TimedTaskResponse, error) {
	var resp pb.TimedTaskResponse
	resp = pb.TimedTaskResponse{
		// ErrCode:
		// EreMsg:
		// ExpData:
	}
	return &resp, nil
}

// TimedTaskTime implements Service.
func (s timserviceService) TimedTaskTime(ctx context.Context, in *pb.TimedTaskRequest) (*pb.TimedTaskResponse, error) {
	var resp pb.TimedTaskResponse
	resp = pb.TimedTaskResponse{
		// ErrCode:
		// EreMsg:
		// ExpData:
	}
	return &resp, nil
}
