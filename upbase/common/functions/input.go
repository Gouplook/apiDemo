/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  input
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:51
 */
package functions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Data string
func (d Data) Form(data interface{}) Data {
	return Data(fmt.Sprint(data))
}
func (d Data) String() string {
	return string(d)
}
func (d Data) Int() int {
	data, _ := strconv.Atoi(string(d))
	return data
}
func (d Data) Int64() int64 {
	data, _ := strconv.ParseInt(string(d), 10, 64)
	return data
}
func (d Data) Uint64() uint64 {
	data, _ := strconv.ParseUint(string(d), 10, 64)
	return data
}
func (d Data) Float64() float64 {
	data, _ := strconv.ParseFloat(string(d), 64)
	return data
}
func (d Data) Bool() bool {
	data, _ := strconv.ParseBool(string(d))
	return data
}
func (d Data) StringArray(sep... string) []string {
	sp := ","
	if len(sep) > 0 {
		sp = sep[0]
	}
	return strings.Split(strings.Trim(string(d), "[] "), sp)
}
func (d Data) IntArray(sep... string) []int {
	sp := ","
	if len(sep) > 0 {
		sp = sep[0]
	}
	data := strings.Split(strings.Trim(string(d), "[] "), sp)
	var output []int
	for _,v := range data {
		output = append(output, Data(v).Int())
	}
	return output
}
type input struct {
	Ctx *gin.Context
	encry bool  // 加密
}
// 判断是否加密
func (i *input)IsEncry() bool{
	return i.encry
}

// 初始化
func (i *input)init(){
	i.encry = false
}

// 获取字符串原始数据
func (i *input)getSrcStr(str string) string {
	if i.IsEncry() {
		return DecodeStr(str)
	}
	return str
}
