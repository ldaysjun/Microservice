package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {


	//res()
	//GetInfo()

	test()
	//fmt.Println(rsp)
}

func res()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		log.Fatal(err)
	}

	rsp,err := cli.Put(context.TODO(),"/services/book/127.0.0.1:8082","127.0.0.1:8080",clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}

	log.Println(rsp)
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


	rsp,err := cli.Get(context.TODO(),"/services/book/",clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	}

	for _,kv := range rsp.Kvs{
		fmt.Println("key=",string(kv.Key),"|","value=",string(kv.Value))
	}
}

func Logout()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})

	defer cli.Close()
	if err != nil {
		log.Fatal(err)
	}

	rsp,err := cli.Delete(context.TODO(),"/services/book/127.0.0.1:2379")
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp.PrevKvs)
}

func test()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "/services/book/127.0.0.1:8088", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// the key 'foo' will be kept forever
	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}

	
	ka := <-ch
	fmt.Println("ttl:", ka.TTL)
}

