/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 10:14
@Description:

*********************************************/
package functions
// 公共方法
// @kancun Team
// @Contact ly@900sui.com

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

//md5加密
//@param  string str 待加密md5字符串
//@return string
func HashMd5(str string) string{
	md5Inst := md5.New()
	md5Inst.Write([]byte(str))
	result := md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x",result)
}
//sha1加密
//@param  string str 待加密sha1字符串
//@return string
func HashSha1(str string) string{
	sha1Inst := sha1.New()
	_,err := sha1Inst.Write([]byte(str))
	if err != nil{
		log.Fatal(err.Error())
	}
	result := sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x",result)
}

//base64加密
//@param  string str 待加密字符串
//@return string
func Base64Encode(str string) string{
	//转换成byte类型
	strB := []byte(str)
	return base64.StdEncoding.EncodeToString(strB)
}

//base64解密
//@param  string str 待解密字符串
//@return string
func Base64Decode(str string) string{
	//转换成byte类型
	bytes,_ := base64.StdEncoding.DecodeString(str)
	return string(bytes[:])
}

//验证手机号
//@param  string phone 待验证手机号
//@return bool
func CheckPhone(phone string) bool{
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

//验证固话
//@param  string tel 待验证固定电话
//@return bool
func CheckCall(tel string) bool{
	//分析参数
	if tel == ""{
		return false
	}
	pattern := "^[\\d]{3,4}\\-[\\d]{7,8}$"
	if bools,_ := regexp.MatchString(pattern,tel);bools{
		return true
	}
	return false
}


//验证邮箱
//@param  string email 待验证邮箱
//@return bool
func CheckEmail(email string) bool{
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//马赛克中国大陆手机号
//@param  string     phone 待打马赛克手机号
//@param  ...string  re    马赛克默认标识 默认="*"
//@return string
func MarkPhone(phone string,re ...string) string{
	if len(phone) != 11 {
		return phone
	}
	var replaceMark string
	if len(re) == 0{
		replaceMark = strings.Repeat("*",5)
	}else{
		replaceMark = strings.Repeat(string(re[0]),5)
	}
	replace := phone[3:8]
	return strings.Replace(phone,replace,replaceMark,1)
}

//使用gob编码将数据转化为byte切片
//@param  interface{} data gob数据
//@return mixted
func GobEncode2Byte(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//gob编码的byte切片数据转化为其他数据
//@param  byte data 字节切片数组
//@return error
func GobDecodeByte(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

//string字符串转json输出
//@param string str 待转字符串
//@return string
func StringsToJSON(str string) string {
	var jsons bytes.Buffer
	for _, r := range str {
		rint := int(r)
		if len(string(r)) == 4 {
			jsons.WriteRune(r)
		} else if rint < 128 {
			jsons.WriteRune(r)
		} else {
			jsons.WriteString("\\u")
			if rint < 0x100 {
				jsons.WriteString("00")
			} else if rint < 0x1000 {
				jsons.WriteString("0")
			}
			jsons.WriteString(strconv.FormatInt(int64(rint), 16))
		}
	}
	return jsons.String()
}


func ParseAcceptLang(ctx *gin.Context) {
	if lang := ctx.Request.Header.Get("Accept-Language"); lang != "" {
		// langs := strings.Split(lang, ":")
	}
}

//把数组转换为字符串
//@param  string separator       转换分隔符
//@param  interface{}  interface 待转换数据
//@return string
func Implode(separator string, array interface{}) (string) {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", separator, -1)
}

//生成指定长度的随机字符串
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
//@param  int n 待生成随机字符串的长度
//@return string
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//字符串转化为时间戳
//@param  string timeStr 日期字符串
//@return int64
func StrtoTime(timeStr string, timelayouts... string) int64 {
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	if len(timelayouts) > 0 {
		timeLayout = timelayouts[0]
	}
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //使用模板在对应时区转化为time.time类型
	return  theTime.Unix()
}

//时间戳转化为字符串
//@param  int64  timestamp  时间戳
//@return string
func TimeToStr(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006/01/02 15:04:05")
}



//获取字符串长度
//@param  string str 待获取长度字符串
//@return int
func Mb4Strlen(str string) int{
	str = strings.TrimSpace(str)
	if len(str) == 0{
		return 0
	}
	strRune := []rune(str)
	lens := len(strRune)
	return lens
}

//截取字符串
//@param string str   待截取的字符串
//@param int    index 截取开始位置
//@param int    lens  截取长度
func StuffStr(str string,index int,lens int)(string){
	str = strings.TrimSpace(str)
	if len(str) == 0{
		return str
	}
	strRune := []rune(str)
	if len(strRune)<lens{
		lens = len(strRune)
	}
	return string(strRune[index:lens])
}



//map转数组
func ArrayKeys (maps map[int]interface{})[]int{
	//分析参数
	if len(maps) == 0{
		return make([]int,0)
	}
	var arr = make([]int,0)
	for i,_ := range maps{
		arr = append(arr,i)
	}
	return arr
}

//map数组转数组
func ArrayValue2Array(field string,maps []map[string]interface{})[]int{
	//分析参数
	if len(maps) == 0{
		return make([]int,0)
	}
	var arr = make([]int,0)
	for _,m := range maps{
		v,ok := m[field]
		if ok{
			if vs,p := v.(string);p{
				n,_ := strconv.Atoi(vs)
				arr = append(arr,n)
			}
			if vs,p := v.(int);p{
				arr = append(arr,vs)
			}
		}
	}
	return arr
}

//map数组转map
func ArrayRebuild(field string,maps []map[string]interface{})map[string]interface{}{
	//分析参数
	if len(maps) == 0{
		return make(map[string]interface{},0)
	}
	var reMap = make(map[string]interface{})
	for _,m := range maps{
		v,ok := m[field]
		if ok{
			if vs,p := v.(int);p{
				reMap[strconv.Itoa(vs)] = m
			}
			if vs,p := v.(string);p{
				reMap[vs] = m
			}
			if vs,p := v.(float64);p{
				reMap[strconv.FormatFloat(vs,'f',-1,64)] = m
			}
			if vs,p := v.(float32);p{
				reMap[strconv.FormatFloat(float64(vs),'f',-1,64)] = m
			}
		}
	}
	return reMap
}


//数组map排序
func SortsMap(field string,maps []map[string]interface{})[]map[string]interface{}{
	var mapData = make(map[string]interface{})
	var keys  = make([]string,0)
	for _,v := range maps{
		vs := v[field]
		if vp,ok := vs.(float64);ok{
			vs = strconv.FormatFloat(vp,'f',-1,64)
		}
		if vp,ok := vs.(int);ok{
			vs = strconv.FormatInt(int64(vp),10)
		}
		if vp,ok := vs.(string);ok{
			vs = vp
		}
		mapData[vs.(string)] = v
		keys = append(keys,vs.(string))
	}
	sort.Strings(keys)
	remapData := make([]map[string]interface{},0)
	for _,v :=range keys{
		remapData = append(remapData,mapData[v].(map[string]interface{}))
	}
	return remapData
}

// InArray
func InArray(search interface{}, array interface{}) bool {
	if arr,ok := array.([]int); ok {
		for _,val := range arr {
			if val == search {
				return true
			}
		}
	}
	if arr,ok := array.([]string); ok {
		for _,val := range arr {
			if val == search {
				return true
			}
		}
	}
	return false
}

//整型数组去重
func ArrayUniqueInt(arr []int)([]int){
	if len(arr) == 0{
		return arr
	}
	newArr := make([]int,0)
	for i:=0;i<len(arr);i++{
		repeat := false
		for j:=i+1;j<len(arr);j++{
			if arr[i] == arr[j]{
				repeat = true
				break
			}
		}
		if arr[i] == 0{
			continue
		}
		if repeat == false{
			newArr = append(newArr,arr[i])
		}
	}
	return newArr
}

//整型数组去重
func ArrayUniqueString(arr []string)([]string){
	if len(arr) == 0{
		return arr
	}
	newArr := make([]string,0)
	for i:=0;i<len(arr);i++{
		repeat := false
		for j:=i+1;j<len(arr);j++{
			if arr[i] == arr[j]{
				repeat = true
				break
			}
		}
		if arr[i] == ""{
			continue
		}
		if repeat == false{
			newArr = append(newArr,arr[i])
		}
	}
	return newArr
}

// ClientIP 尽最大努力实现获取客户端 IP。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// 公钥转换
func GetPemPublic(public_key string) string {
	res := "-----BEGIN PUBLIC KEY-----\n"
	strlen := len(public_key)
	for i:=0;i < strlen;i+=64 {
		if i + 64 >= strlen {
			res += public_key[i:] + "\n"
		}else{
			res += public_key[i:i + 64] + "\n"
		}
	}
	res += "-----END PUBLIC KEY-----"
	return res
}

// 私钥转换
func GetPemPrivate(private_key string) string {
	res := "-----BEGIN RSA PRIVATE KEY-----\n"
	strlen := len(private_key)
	for i:=0;i < strlen;i+=64 {
		if i + 64 >= strlen {
			res += private_key[i:] + "\n"
		}else{
			res += private_key[i:i + 64] + "\n"
		}
	}
	res += "-----END RSA PRIVATE KEY-----"
	return res
}

//字符串切割成int型数组
func StrExplode2IntArr( s string, step string) []int {
	strs := strings.Split(s, ",")
	var outData []int
	for _, v := range strs{
		if len(v) == 0{
			continue
		}
		intv, _ :=strconv.Atoi(v)
		outData = append( outData, intv )
	}
	return outData
}
