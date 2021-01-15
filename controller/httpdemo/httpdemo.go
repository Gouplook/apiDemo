/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  httpdemo
 * @Version: 1.0.0
 * @Date: 2021/1/11 21:50
 */
package httpdemo

import (
	"apiDemo/upapiauth/apiAuth"
	"apiDemo/upbase/common/toolLib"
	"encoding/json"
	"io/ioutil"
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

}
