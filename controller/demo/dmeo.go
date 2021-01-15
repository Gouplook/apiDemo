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
	_const "apiDemo/lang/const"
	"apiDemo/rpcinterface/interface/common"
	"apiDemo/rpcinterface/interface/demo"
	"apiDemo/upapiauth/apiAuth"
	"apiDemo/upbase/common/functions"
	"apiDemo/upbase/common/toolLib"
	"github.com/gin-gonic/gin"
	"strings"
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

// 正常的请求
// bstoken和Utoken说明：
// busController  继承 userController
func (d *DemoController) GetDemo() {

	// 1: 获取输入值
	busId := d.Input.GetPost("busId").Int()
	shopId := d.Input.GetPost("shopId").Int()
	businessYear := d.Input.GetPost("businessYear").Float64()
	isLimit := d.Input.GetPost("isLimit").Bool() // 是否分页
	bindId := d.Input.GetPost("bindId").String() //多个以","号隔开
	companyName := d.Input.GetPost("companyName").String()
	contactCall := d.Input.GetPost("contactCall").String() //手机号码
	businessHoursStart := d.Input.GetPost("businessHoursStart", "").String()
	businessHoursEnd := d.Input.GetPost("businessHoursEnd", "").String()

	lng := d.Public.Lng  // 在header头中设置   参考：//附近门店详情apibus-busfront- GetNearbyShopInfo
	lat := d.Public.Lat
	// 2：需要对前端传过来的数据进行校验
	companyName = strings.TrimSpace(companyName)
	if companyName == "" {
		d.Output.Error(_const.BUSSETTLED_COMPANYNAME_NIL)
		return
	}
	// 限制字数
	if functions.Mb4Strlen(companyName) > 30 {
		d.Output.Error(_const.PARAMS_ERR)
		return
	}
	//手机号码验证
	if len(contactCall) == 11 {
		if !functions.CheckPhone(contactCall) { // 检测若为中国大陆手机号
			d.Output.Error(_const.PARAMS_ERR)
			return
		}
	}else { //检测为中国大陆固定电话
		if functions.CheckCall(contactCall){
			d.Output.Error(_const.PARAMS_ERR)
			return
		}
	}

	// 3:入参数
	args := new(demo.ArgsDemo)
	args.BusId = busId  //
	args.ShopId = shopId
	args.BusinessYear = businessYear
	args.BsToken = common.BsToken{EncodeStr: d.BusInfo.EncodeStr}
	args.Utoken = common.Utoken{UidEncodeStr: d.UserInfo.UidEncodeStr}

	args.Page = d.Input.GetPost("page", "1").Int()
	args.PageSize = d.Input.GetPost("pageSize", "10").Int() // 每页显示10条
	args.IsLimit = isLimit
	args.BindId = bindId
	args.CompanyName = companyName
	args.ContactCall = contactCall
	// 提炼公共参数
	args.Common.BusinessHoursStart = businessHoursStart
	args.Common.BusinessHoursEnd = businessHoursEnd

	args.Lat = lat
	args.Lng = lng

	reply := &demo.ReplyDemo{}

	// 4： 调用逻辑方法Logic方法
	err := new(demo2.DemoLogic).GetDemo(d.Ctx.Request.Context(), args, reply)
	// 5： 返回数据类型到前端
	if err != nil {
		d.Output.Error(toolLib.GetKcErrCode(err), toolLib.GetKcErrMsg(err))
		return
	}

	if args.BusId == 0 {
		d.Output.Error(toolLib.CreateKcErr(_const.PARAMS_ERR))
		return
	}

	if args.ShopId == 0 {
		d.Output.Error(_const.PARAMS_ERR)
		return
	}
	d.Output.Success(map[string]interface{}{"Result": reply}) // 表示reply返回一个ture或false
	d.Output.Success(reply)

}
