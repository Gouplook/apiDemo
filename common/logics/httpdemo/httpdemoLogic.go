/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/15 16:34
@Description:

*********************************************/
package httpdemo

import (
	"apiDemo/rpcinterface/interface/httpdemo"
	httpdemo2 "apiDemo/rpcinterface/client/httpdemo"  // 客户端
	"context"
)

type HttpDemoLogic struct {

}

func (h *HttpDemoLogic)SendDemo(ctx context.Context, args *httpdemo.ArgsHttpDemo, reply *httpdemo.ReplyHttpDemo) error{
	client := new(httpdemo2.HttpDemo).Init()
	defer client.Close()

	if err := client.SendDemo(ctx, args, reply); err != nil {
		return err
	}
	return nil

}


