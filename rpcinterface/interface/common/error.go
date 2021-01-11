/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  error
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:17
 */
package common

import "apiDemo/upbase/common/toolLib"

const (
	// 公共服务验证
	ENCODE_IS_NIL       = "1000001"
	ENCODE_ERR          = "1000002"
	ENCODE_DATA_TIMEOUT = "1000003"
	PERMISSION_ERR      = "1000004"
)


var errMsg = map[string]string{
	// 公共服务验证
	ENCODE_IS_NIL:       "EncodeStr数据为空",
	ENCODE_ERR:          "解密失败",
	ENCODE_DATA_TIMEOUT: "数据已过期",
	PERMISSION_ERR:      "没有操作权限",

}

// 获取错误信息
func GetInterfaceError(code string) error {
	if val, ok := errMsg[code]; ok {
		return toolLib.CreateKcErr(code, val)
	}
	return toolLib.CreateKcErr(code)
}