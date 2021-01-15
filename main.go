/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/1/10 15:00
 */
package main

import (
	"apiDemo/common/component"
	"apiDemo/upbase/common/plugins/jaeger"
	"apiDemo/upgin"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 定义服务
	service := gin.Default()

	//调用logger初始化
	component.InitLogger()

	// 启动链路追踪
	_, closer, err := jaeger.OpenJaeger()
	if err == nil && closer != nil {
		defer closer.Close()
		// 调用中间件
		service.Use(jaeger.SpanMiddle())
	}
	upgin.Bind(service)
	httpHostPort := fmt.Sprintf("%v:%v",upgin.AppConfig.String("httphost"),upgin.AppConfig.String("httpport"))
	err = service.Run(httpHostPort)
	if err != nil {
		log.Fatalf("server start field:%v", err)
	}
}
