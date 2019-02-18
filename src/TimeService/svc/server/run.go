package server

import (
	"TimeService"
	"TimeService/handlers"
	"TimeService/svc"
	"github.com/go-kit/kit/sd/etcdv3"
	"go.etcd.io/etcd/clientv3"
	"time"

	logger "github.com/go-kit/kit/log"


	"context"
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
	go handlers.InterruptHandler(errC)

	go func() {
		listen,err := net.Listen("tcp","localhost:50051")
		if err != nil {
			errC<-err
			return
		}
		srv := MakeGRPCServer(endpoints)
		s := grpc.NewServer()
		echo.RegisterTimServiceServer(s,srv)

		Register()
		//Register_t("127.0.0.1:50051")


		errC<-s.Serve(listen)
	}()

	log.Println("exit",<-errC)
	Exit()


}

func Register()  {

	cli,err := etcdv3.NewClient(context.Background(),[]string{"127.0.0.1:2379"},etcdv3.ClientOptions{
		DialTimeout:time.Second * 3,
		DialKeepAlive:time.Second * 3,
	})
	if err != nil {
		log.Println(err)
	}

	lo := logger.NewNopLogger()
	registrar := etcdv3.NewRegistrar(cli,etcdv3.Service{
		Key:"/Echo/TimService/TimedTask/127.0.0.1:50051",
		Value:"127.0.0.1:50051",
		//TTL:etcdv3.NewTTLOption(time.Second * 3,time.Second * 10),
	},lo)

	if err := lo.Log("msg", "message"); err != nil {
		log.Println(err)
	}

	registrar.Register()

}

func Exit() {
	cli,err := etcdv3.NewClient(context.Background(),[]string{"127.0.0.1:2379"},etcdv3.ClientOptions{
		DialTimeout:time.Second * 3,
		DialKeepAlive:time.Second * 3,
	})
	if err != nil {
		log.Println(err)
	}

	cli.Deregister(etcdv3.Service{
		Key:"/Echo/TimService/TimedTask/127.0.0.1:50051",
		Value:"127.0.0.1:50051",
		//TTL:etcdv3.NewTTLOption(time.Second * 3,time.Second * 10),
	})
	log.Println("exit")
}


func Register_t(addr string)  {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:3 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		log.Fatal(err)
	}

	lease := clientv3.NewLease(cli)
	leaseRsp,err := lease.Grant(context.TODO(),15)
	if err != nil {
		log.Println(err)
	}
	leaseId := leaseRsp.ID
	_,err = cli.Put(context.TODO(),"/Echo/TimService/TimedTask/127.0.0.1:50051",addr,clientv3.WithLease(leaseId))
	if err != nil {
		log.Println(err)
	}

	//频率为ttl 3/1
	leaseChan,err := cli.KeepAlive(context.TODO(),leaseId)
	if err != nil {
		log.Println(err)
	}
	<-leaseChan
	//for {
	//	select {
	//	case leaseKeepResp := <-leaseChan:
	//		if leaseKeepResp == nil {
	//			log.Println("keep alive closed")
	//			return
	//		}else {
	//			log.Println("success",leaseKeepResp.TTL)
	//		}
	//
	//	}
	//}
}




