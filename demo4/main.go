package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		getRes *clientv3.GetResponse
	)

	config = clientv3.Config{
		Endpoints:[]string{"192.168.1.114:2379"}, //集群列表
		DialTimeout:5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	if getRes, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getRes.Kvs)
	}

}

