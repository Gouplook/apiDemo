/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  busController
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:47
 */
package ApiAuth

import (
	"apiDemo/rpcinterface/interface/bus"
	"github.com/gin-gonic/gin"
)

// 总店控制
type BusController struct {
	UserController
	BusInfo bus.ReplyBusAuth
	BusId int
	ShopId  int
}
func (b *BusController)Prepare(ctx *gin.Context) {
	b.UserController
}
