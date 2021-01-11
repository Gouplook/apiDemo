/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  httpdemo
 * @Version: 1.0.0
 * @Date: 2021/1/11 21:50
 */
package httpdemo

import "apiDemo/upapiauth/apiAuth"

type HttpDemoController struct {
	apiAuth.SignController
}

func (b *HttpDemoController)SendDemo(){

}
