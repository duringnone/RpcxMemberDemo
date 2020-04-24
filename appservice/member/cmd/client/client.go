package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcxMember/appservice/member/service"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
	addr0 = flag.String("addr0", "localhost:8973", "server address")
	addr1 = flag.String("addr1", "localhost:8974", "server address")
	forkAddr0 = flag.String("forkAddr0", "localhost:8975", "server address")
	forkAddr1 = flag.String("forkAddr1", "localhost:8976", "server address")
)

func main() {
	//ClientDemo1()
	//ClientDemo2()
	//ClientDemo3()
	ClientFork()
}

func ClientDemo1() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("ServiceUser", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := service.Args{
		Uid: 999,
	}

	reply := &service.Reply{}
	err := xclient.Call(context.Background(), "UserInfo", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(args.Uid, ":", reply.User)
}


func ClientDemo2() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("ServiceUserFn", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := service.Args{
		Uid: 888,
	}

	reply := &service.Reply{}
	err := xclient.Call(context.Background(), "UserReply", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(args.Uid, ":", reply.User)
}

// 点对多
func ClientDemo3() {
	flag.Parse()

	//go ClientP2M(*addr0)
	//go ClientP2M(*addr1)
	//select{}

}
func ClientP2M(addr string){
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	xclient := client.NewXClient("ServiceUser", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := service.Args{
		Uid: 999,
	}

	reply := &service.Reply{}
	err := xclient.Call(context.Background(), "UserInfo", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(args.Uid, ":", reply.User)
}

// 多点发现服务，Fork()
func ClientFork() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *forkAddr0, Value: "weight=1"}, {Key: *forkAddr1, Value: "weight=9"}})	// 设置各个节点的权重，权重高的提供被访问的概念更大
	xclient := client.NewXClient("ServiceUser", client.Failtry, client.WeightedRoundRobin, d, client.DefaultOption)

	defer xclient.Close()

	args := service.Args{
		Uid: 999,
	}

	reply := &service.Reply{}
	err := xclient.Fork(context.Background(), "UserInfo", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(args.Uid, ":", reply.User)
	//m := <- d.WatchService()
	//fmt.Println("当前访问的服务节点：",m)

}