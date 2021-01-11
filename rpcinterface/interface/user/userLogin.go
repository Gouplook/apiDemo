/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  userLogin
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:46
 */
package user

import "context"

type Channel = int // 注册渠道  0=未知， 1=pc网站 2=900岁app 3=康享宝app 4=900岁wap版，5=卡D兜小程序

// 验证登录参数
type CheckLoginParams struct {
	Channel
	Token string // 登录token
}

// 验证登录返回数据
type CheckLoginReply struct {
	UidEncodeStr string // 加密后的uid
	Nick         string // 用户昵称
	RealNameAuth int    // 是否实名认证
}

type UserLogin interface {
	// 验证登录
	CheckLogin(ctx context.Context, args *CheckLoginParams, reply *CheckLoginReply) error

}