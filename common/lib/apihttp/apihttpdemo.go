/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  apihttpdemo
 * @Version: 1.0.0
 * @Date: 2021/1/17 9:48
 */
package apihttp

import (
	"apiDemo/upgin/config/logs"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)



const (
	SHANGHAI_CID = 321 // 上海城市Id

)
var city = make(map[int]map[string]string)
// 调式模式
var Runmode string


type ApiHttpDemo struct {
}

// 发送请求体
type SenRequestBody struct {
	RiskBusId int    `json:"riskBusId"`
	Sign      string `json:"sign"`
}

// 在API端访问远程地址
//ResponseBody 响应体
type ResponseBody struct {
	Error    string      `json:"error"`
	ErrorMsg string      `json:"errorMsg"`
	Total    int         `json:"total"`
	Data     interface{} `json:"data"`
}

// 远程调用API接口获取安全码信息
func (a *ApiHttpDemo)GetSecurityCode(senBody *SenRequestBody,cid int)(responseBody *ResponseBody, err error){
	//
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(senBody)
	// 获取生产或测试 URL
	url := city[cid][Runmode]
	if len(url) == 0 {
		return
	}
	return a.do(requestBody,url)
}

func (a *ApiHttpDemo)do(requestBody *bytes.Buffer,url string)(responseBody *ResponseBody, err error){
	requestBodyStr := fmt.Sprintf("%s", requestBody)
	requestBodyStr = strings.Replace(requestBodyStr, " ", "", -1)
	requestBodyStr = strings.Replace(requestBodyStr, "\n", "", -1)
	request, err := http.NewRequest("GET", url, bytes.NewBufferString(requestBodyStr))
	request.Header.Set("Content-Type", "application/json;charset='utf-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		logs.Info("securityCode response:" + string(body))

		var bodyMap map[string]interface{}
		json.Unmarshal([]byte((body)), &bodyMap)
		fmt.Println(bodyMap)
		if err := json.Unmarshal([]byte(string(body)), &responseBody); err == nil {
			return responseBody, nil
		}
	}
	return responseBody, errors.New("request error:" + response.Status + " URL:" + response.Request.URL.String())
}
