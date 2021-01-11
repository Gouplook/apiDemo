/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 15:42
@Description:

*********************************************/
package demo

import (
	"apiDemo/rpcinterface/client"
	"apiDemo/rpcinterface/interface/demo"
	"context"
)

type Demo struct {
	client.Baseclient
}

func (d *Demo)Init() *Demo{
	d.ServiceName = "rpc_demo"
	d.ServicePath = "Demo"
	return d
}

func(d *Demo)GetDemo(ctx context.Context, args *demo.ArgsDemo, reply *demo.ReplyDemo) error{
	return d.Call(ctx, "GetDemo", args, reply)
}
