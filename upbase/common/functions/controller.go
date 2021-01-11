/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:50
 */
package functions

import (
	"apiDemo/upgin"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	upgin.Controller
	Input  *input
	Output *output
	Public public
}
type public struct {
	Utoken  string  // 登录验证token
	Channel int     // 渠道
	Device  int     // 设备
	Version string  // 版本
	Cid     int     // 城市id
	Lng     float64 // 经度
	Lat     float64 // 维度
}

func (c *Controller) Init(ctx *gin.Context, method string) {
	c.Controller.Init(ctx, method) // 调用upgin/utils/controller.go

	c.Input = MakeInput(c)
	c.Output = MakeOutput(c)

	c.Public.Utoken = c.Input.Header("utoken").String()
	c.Public.Channel = c.Input.Header("channel").Int()
	c.Public.Device = c.Input.Header("device").Int()
	c.Public.Version = c.Input.Header("version").String()
	c.Public.Cid = c.Input.Header("cid").Int()
	c.Public.Lng = c.Input.Header("lng").Float64()
	c.Public.Lat = c.Input.Header("lat").Float64()
}
