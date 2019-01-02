package handlers

import (
	pb "TimeService"
	"context"
	"log"
)

func NewSerVice() pb.TimServiceServer{
	return timServiceServer{}
}

type timServiceServer struct {

}

func (s timServiceServer)TimedTask(ctx context.Context,in *pb.TimedTaskRequest) (*pb.TimedTaskResponse,error) {
	log.Println("req:",in.Start)
	var  resp = pb.TimedTaskResponse{
		EreMsg:"sucess",
		ErrCode:1,
	}
	return &resp,nil
}
