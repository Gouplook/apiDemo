/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  signController
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:48
 */
package apiAuth

import (
	"apiDemo/upbase/common/functions"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

//KEY 密钥
const (
	KEY          = "UYHjLD*u1J8X0CguG&WYp9c5FI79kiaw"
	REQUESTCOUNT = 1
	ONESECOND    = 1    // 限制请求时间间隔（秒）
	REQUESTLIMIT = "on" // off and on
)

var   mu = sync.Mutex{}
var cacheMap = make(map[string]RequestCountTime)

type RequestCountTime struct {
	count int64 // 请求的次数
	lastTime int64 // 最后的访问时间
}

// SignController 验签控制
type SignController struct {
	functions.Controller
	Ignore []string
}
//Prepare
func (s *SignController)Prepare(ctx *gin.Context){
	s.Controller.Prepare(ctx)
	if s.Ctx.IsAborted() || functions.InArray(s.Method, s.Ignore){
		return
	}
	if filterRequest(ctx.Request.RequestURI) {
		s.Output.Error(302, "请求频繁")
		s.Ctx.Abort()
		return
	}

	//设置表单缓存大小
	ctx.Request.ParseMultipartForm(1 << 20)

	requestSign, _ := ctx.GetPostForm("sign")
	sign := encodeSign(ctx.Request.Form)
	if requestSign != sign && os.Getenv("MSF_ENV") != "dev" {
		s.Output.Error(301, "签名参数无效")
		s.Ctx.Abort()
		return
	}
}

// 加码Sign
func encodeSign(postData map[string][]string) (sign string){
	var sslice []string
	postData["key"] = []string{
		KEY,
	}
	for k, _ := range postData {
		sslice = append(sslice, k)
	}
	sort.Strings(sslice[:])

	paramsText := bytes.Buffer{}
	for _, v := range sslice {
		if v == "sign" {
			continue
		}
		paramsText.WriteString(v + "=")
		paramsText.WriteString(postData[v][0] + "&")
	}
	str := strings.Trim(paramsText.String(), "&")
	sign = strings.ToUpper(MD5String(str))

	return
}


// 限制频繁请求
func filterRequest(url string) bool {
	if REQUESTLIMIT == "off" {
		return false
	}
	mu.Lock()
	value, ok := cacheMap[url]
	if ok {
		res := time.Now().Unix() - value.lastTime
		// 最后请求时间和当前时间间隔秒 代表请求间隔太短
		if time.Duration(res) <= ONESECOND {
			mu.Unlock()
			return true
		}

		// 最后请求时间和当前时间间隔秒大于限制就重置当前url最后请求时间
		if time.Duration(res) >= ONESECOND {
			mu.Unlock()
			value.lastTime = time.Now().Unix()
			cacheMap[url] = value
			return false
		}
		value.count++
		cacheMap[url] = value
	} else {
		cacheMap[url] = RequestCountTime{count: 1, lastTime: time.Now().Unix()}
	}

	mu.Unlock()
	return false
}

//MD5String MD5String
func MD5String(str string) (md5Str string) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return
	}
	arr := m.Sum(nil)
	md5Str = fmt.Sprintf("%x", arr)
	return
}