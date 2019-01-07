package svc

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"

	pb "TimeService"
)

type Endpoints struct {
	TimeTaskEndpoint endpoint.Endpoint
}

func MakeTimedTaskEndpoints(s pb.TimServiceServer) endpoint.Endpoint {



	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("cehsi1")
		req := request.(*pb.TimedTaskRequest)
		rsp, err := s.TimedTask(ctx, req)
		if err != nil {
			return nil, err
		}
		return rsp, nil
	}
}

func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"TimedTask": struct{}{},
	}
	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			//log.Println("excluded endpoint ", ex, " does not exist")
		}
		delete(included, ex)
	}

	for incKey := range included {
		fmt.Println(incKey)
		if incKey == "TimedTask" {
			e.TimeTaskEndpoint = middleware(e.TimeTaskEndpoint)
		}
	}

}
