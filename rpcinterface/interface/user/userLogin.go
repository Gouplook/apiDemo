/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  userLogin
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:46
 */
package user

// 验证登录返回数据
type CheckLoginReply struct {
	UidEncodeStr string //加密后的uid
	Nick         string // 用户昵称
	RealNameAuth int    // 是否实名认证
}