# 简介
### 目录简介
```$xslt
└─appservice            // 所有服务总入口目录
   └─member            // 业务模块member
    |   ├─cmd           // cmd目录为client/server入口程序
    |   │  ├─client     // 客户端入口 （模拟客户端用户请求）
    |   │  │      client.go // 
    |   │  │
    |   │  └─server     // 服务端入口（模拟服务端控制器，对client提供服务）
    |   │          server.go
    |   │
    |   ├─model         // model数据层（处理表数据，CURD）
    |   │      member.go
    |   │
    |   └─service           // service服务层（复用逻辑层）
    |           service.go  //(定义了用户服务的所有方法和参数类型结构)
    |            
    |            
    |                 
    └─memberproto        // 业务模块memberproto(测试协议)
    |   ├─cmd           // cmd目录为client/server入口程序
    |   │  ├─client     // 客户端入口 （模拟客户端用户请求）
    |   │  │      client.go // 
    |   │  │
    |   │  └─server     // 服务端入口（模拟服务端控制器，对client提供服务）
    |   │          server.go
    |   │
    |   ├─model         // model数据层（处理表数据，CURD）
    |   │      member.go
    |   │
    |   └─service           // service服务层（复用逻辑层）
    |           service.go  //(定义了用户服务的所有方法和参数类型结构)
    |            
    | 
    └─ membermeta
    |    ├─cmd
    |    │  └─state
    |    │      ├─client
    |    │      │      client.go
    |    │      │
    |    │      └─server
    |    │              server.go
    |    │
    |    ├─model
    |    │      member.go
    |    │
    |    └─service
    |            service.go           
    |                       
    |                       
    |                      
    └─goods        // 业务模块goods  ...........
                ............                    
```


## 笔记
#### Etcd
```
    Etcd 是强一致性的分布式键值数据库
    问题：
        1）键值数据库的应用场景和优/缺点
        2) Etcd如何保证强一致性，Raft协议？proxs协议？？
        3）
        
```

#### 注意项
```
    1）client对象中 4 种调用server方式比较：
        简介：Call()同步(常用); Go() 异步FailMode不生效;Broadcast()和Fork()时，FailMode和selectMode不生效
        A)同步调用：Call()是同步调用server的，可设置FailMode,和 Option,直到server返回成功/错误
        B)异步调用：Go()是异步调用, 异步调用FailMode不生效， 
        C)广播：Broadcast()请求所有server节点，只有所有节点都正常返回才算成功，并取其中一个成功结果返回；若有任一节点返回错误，将返回其中一个错误；失败模式FailMode和路由选择selectMode不生效，建议设置超时，避免client程序挂起
        D)Fork模式：Fork(),将请求发送给所有server节点，任一server节点返回成功则成功；所有节点都返回错误，则选择其中一个错误返回；失败模式FailMode和路由选择selectMode不生效
    2）设置自定义编码/解码器
      A) 详情参见 gobServer.go,gobClient.go,gobService.go
      
    3）超时设置： client/server的保护机制，避免服务无限等待
      A)服务端
        func WithReadTimeout(readTimeout time.Duration) OptionFn
        func WithWriteTimeout(writeTimeout time.Duration) OptionFn
        
      B)客户端
        可在Option.ConnectTimeout,Option.ReadTimeout,Option.WriteTimeout,DefaultOption中默认设置ConnectTimeout=10秒,ReadTimeout,和WriteTimeout,没设置，也可设置context.WithTimeout来控制超时（推荐）
    
   4）metadata元数据
        A) meta中的服务状态
        结论： 服务端可不用设置state=inactive，服务降级仍有效果，但建议还是设置state=inactive。但若客户端不设置这个元数据，则服务端无论怎样设置都不会起作用
       B) 服务分组Group,设置和state类似，客户端必须设置才生效，但服务端建议设置（server不设置也会生效）
        
    
```

#### 安装
```
    1） 问题(protoc-gen-go工具不存在)：--go_out: protoc-gen-go: Plugin failed with status code 1
    解决方案：（安装工具protoc-gen-go） 
        GIT_TAG="v1.2.0" # change as needed
        go get -d -u github.com/golang/protobuf/protoc-gen-go
        git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
        go install github.com/golang/protobuf/protoc-gen-go

```