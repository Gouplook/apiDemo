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

type DemoBase struct {
	Id    int
	Name  string
	Ctime int64
}

type CommonParameter struct {
	BusinessHoursStart string //营业开始时间 格式如：09:00
	BusinessHoursEnd   string //营业结束时间 格式如：22:00
}

type ArgsDemo struct {
	common.BsToken // 店铺信息
	common.Utoken  // 用户信息

	common.Paging      // 带有页码的
	IsLimit       bool //是否分页

	BusId        int
	ShopId       int
	BusinessYear float64 //经营年限
	CompanyName  string  //企业/商户营业执照名称
	BindId       string  //企业/商户所属兼营行业 存在多个多个以","号隔开 如1,2,3
	ContactCall  string  //企业/商户联系电话(手机号或固话)

	Common CommonParameter // 公共参数

	Lng float64 // 经度 附近门店需要经纬度
	Lat float64 // 维度

}

type ReplyDemo struct {
	TotalNum int
	Lists    []DemoBase
}

type Demo interface {
	//
	GetDemo(ctx context.Context, args *ArgsDemo, reply *ReplyDemo) error
}
