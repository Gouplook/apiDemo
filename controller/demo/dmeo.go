/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  dmeo
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:44
 */
package demo

import (
	demo2 "apiDemo/common/logics/demo"
	"apiDemo/rpcinterface/interface/common"
	"apiDemo/rpcinterface/interface/demo"
	"apiDemo/upapiauth/apiAuth"
	"apiDemo/upbase/common/toolLib"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
	apiAuth.BusController
}

// 忽略鉴权验证
func (d *DemoController) Prepare(ctx *gin.Context) {
	d.Ignore = []string{"GetDemo"}
	//d.BusController.Prepare(ctx) 也可以
	d.UserController.Prepare(ctx)

}

// bstoken和Utoken说明：
// busController  继承 userController

func (d *DemoController) GetDemo() {

	busId := d.Input.GetPost("busId").Int()
	shopId := d.Input.GetPost("shopId").Int()

	args := new(demo.ArgsDemo)
	args.BusId = busId
	args.ShopId = shopId
	args.BsToken = common.BsToken{EncodeStr: d.BusInfo.EncodeStr}
	args.Utoken = common.Utoken{UidEncodeStr: d.UserInfo.UidEncodeStr}
	reply := &demo.ReplyDemo{}
	// 调用Logic方法
	err := new(demo2.DemoLogic).GetDemo(d.Ctx.Request.Context(), args, reply)
	if err != nil {
		d.Output.Error(toolLib.GetKcErrCode(err), toolLib.GetKcErrMsg(err))
		return
	}
	d.Output.Success(reply)

}
