/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  busauth
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:09
 */
package bus

import (
	"apiDemo/rpcinterface/interface/common"
	"context"
)

//鉴权 :是指验证用户是否拥有访问系统的权利
//鉴权有四种
//1、HTTP Basic Authentication
//2、session-cookie
//3、Token 验证
//4、OAuth(开放授权)

 // 鉴权用户信息

//saas鉴权输入参数
//BusId-ShopId 至少传一个
type ArgsBusAuth struct {
	common.Utoken  //用户信息
	BusId  int     //企业/商户ID
	ShopId int     //分店ID
	Path   string  //路径
}

//saas鉴权返回参数
type ReplyBusAuth struct {
	EncodeStr  string  //加密字符串
}

type BusAuth interface {
	//saas统一鉴权
	BusAuth(ctx context.Context, args *ArgsBusAuth, reply *ReplyBusAuth) error
}
