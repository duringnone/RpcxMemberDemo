//package server
package main

import (
	"flag"
	"fmt"
	"rpcxMember/appservice/member/service"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
	addr0 = flag.String("addr0", "localhost:8973", "server address")
	addr1 = flag.String("addr1", "localhost:8974", "server address")
)

func main() {

	//ServerP2PDemo()
	ServerP2MDemo()

}

func ServerP2PDemo() {
	flag.Parse()
	// go func + select 实现协程启动，并无限执行（达到守护进程效果）
	// 点对多
	go ServerP2P()
	select {}
}

// 点对点
func ServerP2P() {
	s := server.NewServer()
	//s.Register(new(service.ServiceUser), "")
	s.RegisterName("ServiceUser", new(service.ServiceUser), "")
	s.RegisterFunction("ServiceUserFn", service.UserReply, "")	// 注册服务方法
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("server run successful !")
}


func ServerP2MDemo() {
	flag.Parse()
	// 点对多
	go CreateServer(*addr0)
	go CreateServer(*addr1)
	select {}
}

// 点对多
func CreateServer(addr string) {
	s := server.NewServer()

	s.RegisterName("ServiceUser", new(service.ServiceUser), "")
	err := s.Serve("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr +"---->server run successful !")
}