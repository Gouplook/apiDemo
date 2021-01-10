/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  userController
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:48
 */
package ApiAuth

import (
	"apiDemo/rpcinterface/interface/user"
	"apiDemo/upbase/common/functions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Ignore   []string
	UserInfo user.CheckLoginReply  //验证登录返回数据
	functions.Controller
}
func(u *UserController)Prepare(ctx *gin.Context){
	u.Controller
}
