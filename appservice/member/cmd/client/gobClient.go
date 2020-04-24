package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"github.com/smallnest/rpcx/protocol"
	"log"
	"rpcxMember/appservice/member/service"
	"rpcxMember/appservice/memberproto/pb"
)

var (
	gobAddr     = flag.String("gobAddr", "localhost:8972", "server address")
)
func main() {
	flag.Parse()

	// register customized codec
	share.Codecs[protocol.SerializeType(5)] = &service.GobCodec{} // 自定义 编/解码器
	option := client.DefaultOption
	option.SerializeType = protocol.SerializeType(5) //指定自定义 编/解码器

	d := client.NewPeer2PeerDiscovery("tcp@"+*gobAddr, "")
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