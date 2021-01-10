/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  utoken
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:14
 */
package common

type Utoken struct {
	UidEncodeStr string //uid加密字符串
}

//获取用户UID
func (u *Utoken) GetUid() (int, error) {

}
