package service

import (
	"fmt"
	"context"
	"rpcxMember/appservice/memberproto/pb"
)

/****************** proto协议测试 *****************/
func (s *ServiceUser) UserInfo2(ctx context.Context, args *pb.Args, reply *pb.Reply) error {
	fmt.Println("service:", args.Id)
	reply.AddTime = 14990093
	reply.Uface = "http://image.xxxx.xxx/t.gif"
	reply.UserName = "Joke"
	reply.UserType = 2
	return nil
}

func (t *ServiceUser) Mul(ctx context.Context, args *pb.ProtoArgs, reply *pb.ProtoReply) error {
	reply.C = args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}