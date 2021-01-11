/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 15:58
@Description:

*********************************************/

//适用于企业/商户ID、门店ID的解密
//针对企业/商户、门店增删改操作必须通过加密字符串传递

package common

import (
	"apiDemo/upbase/common/functions"
	"apiDemo/upgin"
	"encoding/json"
	"github.com/wenzhenxi/gorsa"
	"strconv"
	"strings"
	"time"
)

type BsToken struct {
	EncodeStr string //加密字符串
	busId     int    //解密后的busid
	shopId    int    //解密后的shopId
	busAcc    bool   //是否为拥有店铺权限
	isDecrypt bool   //是否解密
}

//解析结构体
type ReplyDeBusAuthBody struct {
	ShopId int
	BusId  int
	BusAcc bool // 是否拥有店铺权限
}

//解密过程
func (b *BsToken)authDecrypt()error{
	if b.isDecrypt {
		return nil
	}
	if b.EncodeStr == "" {
		err := GetInterfaceError(ENCODE_IS_NIL)
		return err
	}
	//公钥转换
	publickey := functions.GetPemPublic(upgin.AppConfig.String("bstoken.public_key"))
	decodeStr,err := gorsa.PublicDecrypt(b.EncodeStr,publickey)
	if err != nil {
		return GetInterfaceError(ENCODE_ERR)
	}
	decodeArr := strings.Split(decodeStr, "|")
	nowTime := time.Now().Local().Unix()
	expTime, _ := strconv.ParseInt(decodeArr[2], 10, 64)
	if expTime < nowTime {
		//已过期
		return GetInterfaceError(ENCODE_DATA_TIMEOUT)
	}
	decryptStr := decodeArr[1]
	var decryBody ReplyDeBusAuthBody
	json.Unmarshal([]byte(decryptStr),&decryBody)
	b.busId = decryBody.BusId
	b.shopId = decryBody.ShopId
	b.busAcc = decryBody.BusAcc
	b.isDecrypt = true
	return nil
}

//获取企业/商户ID
func (b *BsToken)GetBusId()(int,error) {
	err := b.authDecrypt()
	if err != nil {
		return 0, err
	}
	return b.busId, nil
}

//获取是否允许操作总店铺
func (b *BsToken) GetShopId() (int, error) {
	err := b.authDecrypt()
	if err != nil {
		return 0, err
	}
	return b.shopId, nil
}

//获取是否拥有总店操作权限
func (b *BsToken) GetBusAcc() (bool, error) {
	err := b.authDecrypt()
	if err != nil {
		return false, err
	}
	if b.busAcc == false {
		return false, GetInterfaceError(PERMISSION_ERR)
	}
	return true, nil
}
