/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  encry
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:01
 */
package functions

import (
	"crypto/md5" // 实现了MD5哈希算法
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"
)

var (
	key = getKey()
)

func getKey() []byte {
	m := md5.New()
	m.Write([]byte("yOp81Ll90Ghgc"))
	return []byte(hex.EncodeToString(m.Sum(nil)))
}

// 解码
func EncodeStr(str string) string {
	data := []byte(str)
	dataLen := len(data)
	for k, v := range data {
		data[k] = v ^ key[dataLen%32]
	}
	return base64.StdEncoding.EncodeToString([]byte(GetRandomString(3) + string(data)))
}

// 编码过程
func DecodeStr(str string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	data := decodeBytes[3:]
	dataLen := len(data)
	for k, v := range data {
		data[k] = v ^ key[dataLen%32]
	}
	return string(data)
}
// 获取随机字符串
func GetRandomString(length int, padding ...string) string {
	str := "0123456789abcdef"
	if padding != nil {
		str = padding[0]
	}
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
