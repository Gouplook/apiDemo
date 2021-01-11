/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 13:44
@Description:

*********************************************/
package bus

import (
	"apiDemo/rpcinterface/client"
	"apiDemo/rpcinterface/interface/bus"
	"context"
)

type BusAuth struct {
	client.Baseclient // 继承baseclient中的baseclient
}

//初始化
func (b *BusAuth) Init() *BusAuth {
	b.ServicePath = "BusAuth"
	b.ServiceName = "rpc_bus"
	return b
}

// SAAS 统一鉴权
func (b *BusAuth) BusAuth(ctx context.Context,args *bus.ArgsBusAuth,reply *bus.ReplyBusAuth) error{
	return b.Call(ctx, "BusAuth", args, reply)
}