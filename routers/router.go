/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/1/10 15:11
 */
package routers

import (
	"apiDemo/controller/demo"
	"apiDemo/controller/http"
	"apiDemo/upgin"
)

func init(){
	version := "v1"

	v1 := upgin.Group("/" + version)
	{
		v1.AutoRouter("demo",&demo.DemoController{},"post,get")
		v1.AutoRouter("http", &http.HttpDemoController{},"post,get")
	}

	// 签名授权接口
}