/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  upError
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:19
 */
package toolLib

import (
	"fmt"
	"strings"
)

//根据errcode 返回一个具体的错误信息
//@param errCode string 错误码
//@param errMsg ...string 错误信息
func CreateKcErr( errCode string, errMsg ...string )  error {
	errmsg := ""
	if len( errMsg ) > 0 {
		errmsg = errMsg[0]
	}else{
		errmsg = lang.GetLang( errCode )
	}
	return errors.New(fmt.Sprintf("%s#C#%s", errmsg, errCode ))

}

//返回错误的错误信息
//@param err srror 错误
func GetKcErrMsg( err error ) ( errMsg string )  {
	errs := strings.Split( fmt.Sprint( err ), "#C#" )
	errMsg = errs[0]
	return
}

//返回错误的错误码
//@param err srror 错误
func GetKcErrCode(err error) (errCode string) {
	errCode = ""
	errs := strings.Split( fmt.Sprint( err ), "#C#" )
	if len(errs) == 2{
		errCode = errs[1]
	}
	return
}