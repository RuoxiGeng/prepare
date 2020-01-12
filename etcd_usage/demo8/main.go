package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putop  clientv3.Op
		getop  clientv3.Op
		opResp clientv3.OpResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	//创建Op
	putop = clientv3.OpPut("/cron/jobs/job8", "123")

	//执行Op
	if opResp, err = kv.Do(context.TODO(), putop); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入Revision:", opResp.Put().Header.Revision)

	getop = clientv3.OpGet("/cron/jobs/job8")

	if opResp, err = kv.Do(context.TODO(), getop); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Revision:", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("Value:", string(opResp.Get().Kvs[0].Value))
}
