package server

import (
	"TimeService"
	"TimeService/svc"
	"context"
	"github.com/go-kit/kit/transport/grpc"
)

func MakeGRPCServer(endpoints svc.Endpoints) echo.TimServiceServer{
	return &grpcServer{
		timedtask:grpc.NewServer(
			endpoints.TimeTaskEndpoint,
			DecodeGRPCTimeTaskRequest,
			EncodeGRPCTimeTaskReponse,
			),
	}
}

type grpcServer struct {
	timedtask grpc.Handler
}

func (s *grpcServer)TimedTask(ctx context.Context,req *echo.TimedTaskRequest)(*echo.TimedTaskResponse,error)  {
	_,rep,err := s.timedtask.ServeGRPC(ctx,req)
	if err != nil {
		return  nil,err
	}
	return rep.(*echo.TimedTaskResponse),nil
}

func DecodeGRPCTimeTaskRequest(_ context.Context, grpcReq interface{})(interface{},error)  {
	req := grpcReq.(*echo.TimedTaskRequest)
	return  req,nil
}

func EncodeGRPCTimeTaskReponse(_ context.Context, grpcRsp interface{})(interface{},error)  {
	resp := grpcRsp.(*echo.TimedTaskResponse)
	return resp,nil
}


