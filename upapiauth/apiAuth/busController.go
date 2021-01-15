/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  busController
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:47
 */
package apiAuth

import (
	bus2 "apiDemo/rpcinterface/client/bus"
	"apiDemo/rpcinterface/interface/bus"
	"apiDemo/rpcinterface/interface/common"
	"apiDemo/upbase/common/functions"
	"apiDemo/upbase/common/toolLib"
	"github.com/gin-gonic/gin"
)

// 总店控制
type BusController struct {
	UserController
	BusInfo bus.ReplyBusAuth
	BusId   int
	ShopId  int
}

// 忽略鉴权
func (b *BusController) Prepare(ctx *gin.Context) {
	b.UserController.Prepare(ctx)
	// 上下文是否中止
	if !b.Ctx.IsAborted(){
		b.BusId = b.Input.GetPost("busId").Int()
		b.ShopId = b.Input.GetPost("shopId").Int()
	}
	// 如果没有忽略鉴权的方法，调用鉴权方法获取加密字符串
	if !functions.InArray(b.Method, b.Ignore) {
		// 调用客户端
		busAuth := new(bus2.BusAuth).Init()
		defer busAuth.Close()  // 实际这个close没有实现
		err := busAuth.BusAuth(b.Ctx.Request.Context(),&bus.ArgsBusAuth{
			Utoken:common.Utoken{UidEncodeStr: b.UserInfo.UidEncodeStr},
			BusId:  b.BusId,
			ShopId: b.ShopId,
			Path: b.Ctx.FullPath(),
		},&b.BusInfo)
		if err != nil {
			b.Output.Error(toolLib.GetKcErrCode(err),toolLib.GetKcErrMsg(err))
			b.Ctx.Abort() // 上下文中止
			return
		}
	}
}





