/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  userController
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:48
 */
package apiAuth

import (
	user2 "apiDemo/rpcinterface/client/user"
	"apiDemo/rpcinterface/interface/user"
	"apiDemo/upbase/common/functions"
	"apiDemo/upbase/common/toolLib"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Ignore   []string
	UserInfo user.CheckLoginReply  //验证登录返回数据
	functions.Controller  // 继承base中controller
}
func(u *UserController)Prepare(ctx *gin.Context){
	u.Controller.Prepare(ctx)
	// 判断ctx上下文是否中止 或或被忽略鉴权，直接登陆
	if !u.Ctx.IsAborted() && !functions.InArray(u.Method, u.Ignore) {
		if u.Public.Utoken == "" {
			u.Output.Error("110021", "未登录")
			u.Ctx.Abort()
			return
		}
		// 调用登陆接口
		rpcUserlogin := new(user2.UserLogin).Init()
		defer rpcUserlogin.Close()
		if err := rpcUserlogin.CheckLogin(u.Ctx.Request.Context(), &user.CheckLoginParams{
			Channel: u.Public.Channel,
			Token:   u.Public.Utoken,
		}, &u.UserInfo); err != nil {
			u.Output.Error(toolLib.GetKcErrCode(err), toolLib.GetKcErrMsg(err))
			u.Ctx.Abort()
			return
		}
	}
}
