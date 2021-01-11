/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 15:40
@Description:

*********************************************/
package demo

import (
	"apiDemo/rpcinterface/interface/common"
	"context"
)

type ArgsDemo struct {
	common.BsToken // 店铺信息
	common.Utoken  // 用户信息
	BusId          int
	ShopId         int
}

type ReplyDemo struct {
}

type Demo interface {
	//
	GetDemo(ctx context.Context, args *ArgsDemo, reply *ReplyDemo) error
}
