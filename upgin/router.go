/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/1/10 19:27
 */
package upgin

import (
	"apiDemo/upgin/config/logs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)
const (
	FILTER_DENY = "deny"
	FILTER_ACCESS = "access"
)
type ControllerInfo struct {
	pattern        string // 方式
	controllerName string
	methods        map[string]string
	controllerType reflect.Type
	handler        http.Handler
	initialize     func() ControllerInterface // 接口

}

var (
	controllerInfos = map[string]*ControllerInfo{}
	static           = map[string]string{} // 静态资源
	HTTPMETHOD      = map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"PATCH":   true,
		"OPTIONS": true,
		"HEAD":    true,
		"TRACE":   true,
		"CONNECT": true,
	}
)

type RouterInterface interface {
	Get(pattern string, c ControllerInterface, method string)
	Post(pattern string, c ControllerInterface, method string)
	Put(pattern string, c ControllerInterface, method string)
	Patch(pattern string, c ControllerInterface, method string)
	Head(pattern string, c ControllerInterface, method string)
	Options(pattern string, c ControllerInterface, method string)
	Delete(pattern string, c ControllerInterface, method string)
	Connect(pattern string, c ControllerInterface, method string)
	Trace(pattern string, c ControllerInterface, method string)
	Any(pattern string, c ControllerInterface, method string)
	Router(pattern string, c ControllerInterface, mappingMethods ...string)
}

type routerGroup struct {
	group string
}

func (r *routerGroup) Get(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Get:"+method)
}
func (r *routerGroup) Post(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Post:"+method)
}
func (r *routerGroup) Put(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Put:"+method)
}
func (r *routerGroup) Patch(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Patch:"+method)
}
func (r *routerGroup) Head(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Head:"+method)
}
func (r *routerGroup) Options(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Options:"+method)
}

func (r *routerGroup) Delete(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Delete:"+method)
}

func (r *routerGroup) Connect(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Connect:"+method)
}

func (r *routerGroup) Trace(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "Trace:"+method)
}

func (r *routerGroup) Any(pattern string, c ControllerInterface, method string) {
	addWithMethodParams(r.group+pattern, c, "*:"+method)
}

func (r *routerGroup) Router(pattern string, c ControllerInterface, methods string) {
	addWithMethodParams(r.group+pattern, c, methods)
}

// 控制器自动路由
// extends 拓展字段 0:注册路由的类型-any 1:过滤方法 2:路由名称规则-首字母小写
func (r *routerGroup) AutoRouter(pattern string, c ControllerInterface, extends ...string) {
	autoRouterController(r.group+pattern, c, extends...)
}
func autoRouterController(pattern string, c ControllerInterface, extends ...string) {
	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	methodType := []string{"*"}
	filterType := 0
	filterMethods := map[string]bool{}
	for index,extend := range extends {
		if index == 0 && extend != "" {
			methodType = strings.Split(extend, ",")
		}
		if index == 1 && extend != "" {
			colon := strings.Split(extend, ":")
			filters := []string{}
			if len(colon) == 2 {
				if colon[0] == FILTER_ACCESS {
					filterType = 1
				}
				filters = strings.Split(colon[1], ",")
			}else{
				filters = strings.Split(extend, ",")
			}
			for _,filter := range filters {
				filterMethods[filter] = true
			}
		}
	}

	reflectTof := reflect.TypeOf(c)
	methodNum := reflectVal.NumMethod()
	for index :=0;index < methodNum;index++ {
		method := reflectTof.Method(index)
		if _,ok := NotAutoRouter[method.Name];ok != true {
			methods := make(map[string]string)
			if filterType == 0 && filterMethods[method.Name] {
				continue
			}else if filterType == 1 && !filterMethods[method.Name] {
				continue
			}
			for _,v := range methodType {
				if v == "*" || HTTPMETHOD[strings.ToUpper(v)] {
					methods[strings.ToUpper(v)] = method.Name
				}else{
					panic("'" + v + "' method doesn't exist in the controller " + t.Name())
				}
			}
			addToControllerInfos(fmt.Sprintf("%v/%v",pattern,strFirstToLower(method.Name)), methods, c, t)
		}
	}
}
func strFirstToLower(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 65 && strArry[0] <= 90  {
		strArry[0] +=  32
	}
	return string(strArry)
}
func addWithMethodParams(pattern string, c ControllerInterface, mappingMethods string) {
	if mappingMethods == "" {
		panic("reg router master input method:" + pattern)
	}
	reflectVal := reflect.ValueOf(c)
	// 返回v指向的值
	t := reflect.Indirect(reflectVal).Type()
	methods := make(map[string]string)

	semi := strings.Split(mappingMethods, ";")
	for _, v := range semi {
		colon := strings.Split(v, ":")
		if len(colon) != 2 {
			panic("method mapping format is invalid")
		}
		comma := strings.Split(colon[0], ",")
		for _, m := range comma {
			if m == "*" || HTTPMETHOD[strings.ToUpper(m)] {
				if val := reflectVal.MethodByName(colon[1]); val.IsValid() {
					methods[strings.ToUpper(m)] = colon[1]
				} else {
					panic("'" + colon[1] + "' method doesn't exist in the controller " + t.Name())
				}
			} else {
				panic(v + " is an invalid method mapping. Method doesn't exist " + m)
			}
		}
	}
	if len(methods) == 0 {
		panic("no method reg route:" + pattern)
	}
	addToControllerInfos(pattern, methods, c, t)
}
func addToControllerInfos(pattern string, methods map[string]string, c ControllerInterface, t reflect.Type) {
	pattern = filepath.ToSlash(filepath.Join(pattern))
	route := &ControllerInfo{}
	route.controllerName = t.Name()
	route.pattern = pattern
	route.methods = methods
	route.controllerType = t
	route.initialize = func() ControllerInterface {
		vc := reflect.New(route.controllerType)
		execController, ok := vc.Interface().(ControllerInterface)
		if !ok {
			panic("controller is not ControllerInterface")
		}

		elemVal := reflect.ValueOf(c).Elem()
		elemType := reflect.TypeOf(c).Elem()
		execElem := reflect.ValueOf(execController).Elem()

		numOfFields := elemVal.NumField()
		for i := 0; i < numOfFields; i++ {
			fieldType := elemType.Field(i)
			elemField := execElem.FieldByName(fieldType.Name)
			if elemField.CanSet() {
				fieldVal := elemVal.Field(i)
				elemField.Set(fieldVal)
			}
		}

		return execController
	}
	controllerInfos[pattern] = route
}
func execHandel(ctx *gin.Context)  {
	if ex,ok := controllerInfos[ctx.FullPath()];ok == true {
		exController := ex.initialize()
		vc := reflect.ValueOf(exController)

		methodName := ""
		if method,ok := ex.methods[ctx.Request.Method]; ok == true {
			methodName = method
		} else {
			if method,ok := ex.methods["*"]; ok == true {
				methodName = method
			}
		}
		method := vc.MethodByName(methodName)
		exController.Init(ctx, methodName)
		if !ctx.IsAborted() {
			exController.Prepare(ctx)
		}
		if !ctx.IsAborted() {
			vals := make([]reflect.Value,0)
			method.Call(vals)
		}
		if !ctx.IsAborted() {
			exController.Finish(ctx)
		}
	}
}
// Engine引擎是框架的实例
func Bind(e *gin.Engine) {
	for pattern,cInfo := range controllerInfos {
		for method := range cInfo.methods {
			if method == "*" {
				e.Any(pattern, execHandel)
			} else {
				e.Handle(method, pattern, execHandel)
			}
		}
	}
	for pattern, root := range static {
		e.Static(pattern, root)
	}
}


func GetURL(endpoint string,fields ...interface{}) string {
	controller := path.Base(endpoint)
	paths := strings.Split(controller, ".")
	if len(paths) <= 1 {
		logs.Warn("urlfor endpoint must like path.controller.method")
		return ""
	}
	if len(fields)%2 != 0 {
		logs.Warn("urlfor params must key-value pair")
		return ""
	}
	pkg := ""
	if controller != endpoint {
		pkg = path.Dir(endpoint)
	}
	return getURL(pkg, paths[0], paths[1], fields...)
}

func getURL(pkg string, controllerName string, methodName string, fields ...interface{}) string {
	fieldLen := len(fields)
	if fieldLen % 2 != 0 {
		logs.Warn("urlfor params must key-value pair")
		return ""
	}

	for pattern,cInfo := range controllerInfos {
		if cInfo.controllerName == controllerName {
			if pkg != "" && !strings.HasSuffix(cInfo.controllerType.PkgPath(), pkg) {
				continue
			}
			for method,mName := range cInfo.methods {
				if method == "GET" && mName == methodName {
					flag := true
					for i:=0; i< fieldLen; i+=2 {
						fieldKey := fmt.Sprint(fields[i])
						fieldValue := fmt.Sprint(fields[i+1])
						if strings.HasPrefix(fieldKey, ":") {
							pattern = strings.Replace(pattern, fieldKey, url.QueryEscape(fieldValue), 1)
						}else{
							if flag {
								pattern += "?"
							}else{
								pattern += "&"
							}
							pattern += fmt.Sprintf("%s=%s", url.QueryEscape(fieldKey), url.QueryEscape(fieldValue))
							flag = false
						}
					}
					return pattern
				}
			}
		}
	}
	return ""
}
