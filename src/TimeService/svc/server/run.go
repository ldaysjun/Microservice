package server

import (
	"TimeService"
	"TimeService/handlers"
	"TimeService/svc"

	"google.golang.org/grpc"
	"log"
	"net"
)

type Config struct {
	HTTPAddr string
	DebugAddr string
	GRPCAddr string
}

func NewEndpoints()svc.Endpoints {

	var service echo.TimServiceServer
	{
		service = handlers.NewSerVice()
	}
	var (
		timedTaskEndpoints = svc.MakeTimedTaskEndpoints(service)
	)
	endpoints := svc.Endpoints{
		TimeTaskEndpoint:timedTaskEndpoints,
	}

	endpoints = handlers.WrapEnpoints(endpoints)
	return endpoints
}

func Run()  {
	errC := make(chan error)

	endpoints := NewEndpoints()

	go func() {
		listen,err := net.Listen("tcp","localhost:50051")
		if err != nil {
			errC<-err
			return
		}

		srv := MakeGRPCServer(endpoints)
		s := grpc.NewServer()
		echo.RegisterTimServiceServer(s,srv)
		errC<-s.Serve(listen)
	}()


	log.Println("exit",<-errC)

}

