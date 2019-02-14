package main

import (
	"TimeService/svc/server"
	"context"
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	"log"

	kitlog "github.com/go-kit/kit/log"

	"net"
	"time"
)

func main()  {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	
	//log.Println(GetLocalIP())
	//re()
	
	server.Run()
}

func GetLocalIP() (string,error){
	addrs,err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return "",err
	}

	for _,addr:=range addrs{
		if ipnet,ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(),nil
			}
		}
	}
	panic("unable to determine locla ip")
}

func re()  {
	opertion := etcdv3.ClientOptions{
		DialKeepAlive:time.Second * 3,
		DialTimeout:time.Second * 3,
	}
	client, err := etcdv3.NewClient(context.Background(),[]string{"127.0.0.1:3333"},opertion)
	if err != nil {
		fmt.Println(err)
	}
	res := etcdv3.NewRegistrar(client,etcdv3.Service{
		Key:"test",
		Value:"test",
	},kitlog.NewNopLogger())
	res.Register()


}


