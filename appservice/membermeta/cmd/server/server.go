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

	ServerP2PDemo()

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
	s.RegisterName("ServiceUser", new(service.ServiceUser), "state=inactive") // 设置metadata（元数据），state可用于服务降级（设置服务存活，但一段时间内不可用）
	s.RegisterFunction("ServiceUserFn", service.UserReply, "")	// 注册服务方法
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("server run successful !")
}

