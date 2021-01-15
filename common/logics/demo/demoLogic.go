/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 16:29
@Description:

*********************************************/
package demo

import (
	demo2 "apiDemo/rpcinterface/client/demo"
	"apiDemo/rpcinterface/interface/demo"
	"context"
)

type DemoLogic struct {
}


func (d *DemoLogic) GetDemo(ctx context.Context, args *demo.ArgsDemo, reply *demo.ReplyDemo) (err error) {
	demoClient := new(demo2.Demo).Init()
	defer demoClient.Close()
	if err = demoClient.GetDemo(ctx, args, reply); err != nil {
		return
	}
	return
}
