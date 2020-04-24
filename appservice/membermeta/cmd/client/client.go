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
	addr1 = flag.String("addr1", "localhost:8973", "server address")
)

func main() {
	ClientDemo1()
}

func ClientDemo1() {
	flag.Parse()

	//d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr, Value: "state=inactive"}, {Key: *addr1, Value: "state=inactive"}})
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

