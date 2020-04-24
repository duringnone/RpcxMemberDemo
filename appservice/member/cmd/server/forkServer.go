//package server
package main

import (
	"flag"
	"fmt"
	"rpcxMember/appservice/member/service"
	"github.com/smallnest/rpcx/server"
)

var (
	//addr = flag.String("addr", "localhost:8972", "server address")
	forkAddr0 = flag.String("forkAddr0", "localhost:8975", "server address")
	forkAddr1 = flag.String("forkAddr1", "localhost:8976", "server address")
)

func main() {

	forkServerDemo()

}

func forkServerDemo() {
	flag.Parse()
	// 点对多
	go forkServer(*forkAddr0)
	go forkServer(*forkAddr1)
	select {}
}

// 点对多
func forkServer(addr string) {
	s := server.NewServer()

	s.RegisterName("ServiceUser", new(service.ServiceUser), "")
	err := s.Serve("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr +"---->server run successful !")
}