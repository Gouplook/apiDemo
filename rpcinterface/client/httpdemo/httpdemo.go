/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/15 16:20
@Description:

*********************************************/
package httpdemo

import (
	"apiDemo/rpcinterface/client"
	"apiDemo/rpcinterface/interface/httpdemo"
	"context"
)

type HttpDemo struct {
	client.Baseclient

}


func (h *HttpDemo)Init() *HttpDemo{
	h.ServiceName = "rpc_httpdemo"
	h.ServicePath = "Httpdemo"

	return h
}

func (h *HttpDemo) SendDemo(ctx context.Context, args *httpdemo.ArgsHttpDemo, reply *httpdemo.ReplyHttpDemo) error {
	return h.Call(ctx,"SendDemo", args,reply)
}
