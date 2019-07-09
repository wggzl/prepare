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
		lease clientv3.Lease
		leaseGrandRes *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		kv clientv3.KV
		putRes *clientv3.PutResponse
		getRes *clientv3.GetResponse
		keepRes *clientv3.LeaseKeepAliveResponse
		keepResChan <-chan *clientv3.LeaseKeepAliveResponse
	)

	config = clientv3.Config{
		Endpoints:[]string{"192.168.1.114:2379"}, //集群列表
		DialTimeout:5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//申请一个lease(租约）
	lease = clientv3.NewLease(client)

	//申请一个10秒的租约
	if leaseGrandRes, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}




	//拿到租约的ID
	leaseId = leaseGrandRes.ID

	//自动续租
	//ctx, _ := context.WithTimeout(context.TODO(), 5 * time.Second)

	//续租了5秒 停止了续租 10秒的生命期 =  15秒的生命期

	//5秒后取消续租
	if keepResChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	//处理续约应答的协程
	go func() {
		for {
			select {
				case keepRes = <- keepResChan:
					if keepResChan == nil {
						fmt.Println("租约已经失效")
						goto END
					} else {//每一秒会续租一次 所以就会收到一次应答
						fmt.Println("收到自动续租应答", keepRes.ID)
					}
			}
		}
		END:
	}()

	//获得kv API子集
	kv = clientv3.NewKV(client)

	//put一个KV，让它与租约关联起来 从10秒后自动过期
	if putRes, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入成功：", putRes.Header.Revision)

	//定时查看一下key过期了没有
	for {
		if getRes, err =kv.Get(context.TODO(), "/cron/lock/job1"); err != nil{
			fmt.Println(err)
			return
		}

		if getRes.Count == 0 {
			fmt.Println("kv过期了")
			break
		}

		fmt.Println("还没过期", getRes.Kvs)
		time.Sleep(2 * time.Second)
	}



}

