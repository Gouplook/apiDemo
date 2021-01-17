/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  http
 * @Version: 1.0.0
 * @Date: 2021/1/11 21:50
 */
package http

import (
	"apiDemo/rpcinterface/interface/common"
	"apiDemo/rpcinterface/interface/httpdemo"
	"apiDemo/upapiauth/apiAuth"
	"apiDemo/upbase/common/toolLib"
	"encoding/json"
	"io/ioutil"
	httplogic "apiDemo/common/logics/httpdemo"
)

type HttpDemoController struct {
	apiAuth.SignController
}

// 远程访问http 方式有两种
// 1：在API端实现
// 2：在rpc端实现
func (h *HttpDemoController) SendDemo() {
	// 请求数据
	var requestMap map[string]interface{}
	if body, err := ioutil.ReadAll(h.Ctx.Request.Body); err != nil {
		json.Unmarshal(body, &requestMap)
	} else {
		h.Output.Error(500, toolLib.GetKcErrMsg(err))
		return
	}
	requestData := &httpdemo.ArgsHttpDemo{}
	requestData.Data = requestMap
	reply := &httpdemo.ReplyHttpDemo{}
	var httpDemoLogic *httplogic.HttpDemoLogic

	err := httpDemoLogic.SendDemo(h.Ctx.Request.Context(),requestData,reply)
	if err != nil {
		h.Output.Error(reply.Code,toolLib.GetKcErrMsg(err))
		return
	}

	if reply.Code != common.CODE_SUCCESS {
		h.Output.Error(reply.Code, reply.Message)
		return
	}

	h.Output.Success(reply.Data)
	return
}
