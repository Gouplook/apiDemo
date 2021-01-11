/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 14:47
@Description:

*********************************************/
package user

import (
	"apiDemo/rpcinterface/client"
	"apiDemo/rpcinterface/interface/user"
	"context"
)

type UserLogin struct {
	client.Baseclient
}

func (u *UserLogin)Init() *UserLogin{
	u.ServiceName = "rpc_user"
	u.ServicePath = "UserLogin"  // 路径从那里获取
	return u
}

//验证登录
func (u *UserLogin) CheckLogin(ctx context.Context, args *user.CheckLoginParams, reply *user.CheckLoginReply) error {
	return u.Call(ctx, "CheckLogin", args, reply)
}