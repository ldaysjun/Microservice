package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		log.Fatal(err)
	}

	rsp,err := cli.Put(context.TODO(),"/test/c","something",clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp)
}

func GetInfo()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})

	defer cli.Close()
	if err != nil {
		log.Fatal(err)
	}

	rsp,err := cli.Get(context.TODO(),"/test/c")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(rsp.Kvs[0].Value))
}

