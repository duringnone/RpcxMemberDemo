package service

import (
	"context"
	"fmt"
	"rpcxMember/appservice/member/model"
)



/**
 * 1）
 * 2）服务方法UserInfo()是可导出的（即首字母大写 外部可调用的，*注：Go中方法首字母大写类似public属性）
 * 3）该方法UserInfo()必须有2个可导出或内建类型的参数
 * 4）第一个参数为context.Context,输入参数Args和 输出参数Reply
 * 5）UserInfo()
 * 6）
 */


type Args struct {
	Uid int
}

type Reply struct {
	model.User
}

// ServiceUser 为外部提供服务
type ServiceUser struct {
}

// UserInfo为外部可调用（首字母大写）
func (s *ServiceUser) UserInfo(ctx context.Context, args *Args, reply *Reply) error {
	fmt.Println("service:", args.Uid)
	reply.User.AddTime = 14990093
	reply.User.Uface = "http://image.xxxx.xxx/t.gif"
	reply.User.UID = int64(args.Uid)
	reply.User.UserName = "Joke"
	reply.User.UserType = 2
	return nil
}

// 添加服务，必须在服务端入口注册服务方法才能生效
func UserReply(ctx context.Context, args *Args, reply *Reply) error {
	reply.User.AddTime = 10000999
	reply.User.Uface = "http://image.xxxx.xxx/reply.gif"
	reply.User.UID = int64(args.Uid)
	reply.User.UserName = "Reply"
	reply.User.UserType = 3
	return nil
}

