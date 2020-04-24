package main

import (
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/serverplugin"
	"github.com/smallnest/rpcx/server"
	"rpcxMember/appservice/member/service"
	"time"
	"log"
)

var (
	etcAdd     = flag.String("etcAdd", "localhost:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "192.168.192.1:2379", "etcd address") // ectd的地址，2379位etcd默认端口
	basePath = flag.String("base", "/rpcx_test", "prefix path")	// 客户端显示： /rpcx_test/ServiceUser/tcp@localhost:8972
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addEtcdRegistryPlugin(s)
	//s.Register(new(service.ServiceUser), "")

	s.RegisterName("ServiceUser", new(service.ServiceUser), "")

	err := s.Serve("tcp", *etcAdd)
	if err != nil {
		panic(err)
	}
}

func addEtcdRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *etcAdd,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
