/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/15 16:16
@Description:

*********************************************/
package httpdemo

import (
	"apiDemo/rpcinterface/interface/common"
	"context"
)

// http 远程调用入参数
type ArgsHttpDemo struct {
	common.Input

}

// http 返回远程调用参数
type ReplyHttpDemo struct {
	common.Output

}

type HttpDemo interface {
	SendDemo(ctx context.Context, args *ArgsHttpDemo, reply *ReplyHttpDemo) error
}
