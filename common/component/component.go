/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  component
 * @Version: 1.0.0
 * @Date: 2021/1/10 15:17
 */

//部件
package component

import (
	"apiDemo/upgin"
	"apiDemo/upgin/config/logs"
	"encoding/json"
	"fmt"
	"time"
)

func InitLogger()(err error) {
	//打印环境变量  upgin 自己封装gin中配置文件
	logs.Info("Environment Variable:MSF_ENV:",upgin.UpConfig.RunMode)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	//开启异步缓存
	logs.Async(1e3)
	fileName := time.Now().Format("20060102") + ".log"
	logConf := make(map[string]interface{})
	logConf= map[string]interface{}{
		"filename": "/opt/logs/" +upgin.AppConfig.String("jaeger.serviceName")+"/" + fileName,
		"maxdays" :1,
	}
	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	// 设置一个新的记录器
	logs.SetLogger(logs.AdapterFile, string(confStr))
	return

}
