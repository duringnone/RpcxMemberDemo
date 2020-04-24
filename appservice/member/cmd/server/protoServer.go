package main

import (
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"rpcxMember/appservice/member/service"
)

var (
	ptotoAddr     = flag.String("ptotoAddr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	// go func + select 实现协程启动，并无限执行（达到守护进程效果）
	// 点对多
	go ServerProto()
	select {}
}

// 点对点
func ServerProto() {
	s := server.NewServer()
	//s.Register(new(service.ServiceUser), "")
	s.RegisterName("ServiceUser", new(service.ServiceUser), "")
	s.RegisterFunction("ServiceUserFn", service.UserReply, "")	// 注册服务方法
	err := s.Serve("tcp", *ptotoAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("server run successful !")
}