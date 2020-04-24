package main

import (
	"flag"
	"log"
	"context"
	"rpcxMember/appservice/memberproto/pb"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	protoAddr     = flag.String("protoAddr", "localhost:8972", "server address")
)
func main() {
	flag.Parse()

	// register customized codec
	option := client.DefaultOption
	option.SerializeType = protocol.ProtoBuffer

	d := client.NewPeer2PeerDiscovery("tcp@"+*protoAddr, "")
	xclient := client.NewXClient("ServiceUser", client.Failover, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args1 := &pb.ProtoArgs{
		A: 10,
		B: 20,
	}

	reply1 := &pb.ProtoReply{}
	err := xclient.Call(context.Background(), "Mul", args1, reply1)

	args := &pb.Args{
		Id: 999,
	}

	reply := &pb.Reply{}

	err = xclient.Call(context.Background(), "UserInfo2", args, reply)

	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args1.A, args1.B, reply1.C)
	log.Println(args.Id, ":", reply)
}