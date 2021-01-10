/**
 * @Author: Yinjinlin
 * @Description:  Jaeger版本的链路追踪
 * @File:  jaeger
 * @Version: 1.0.0
 * @Date: 2021/1/10 17:20
 */
package jaeger

import (
	"apiDemo/upgin"
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"io/ioutil"
)

// 创建追踪者（tracer）
//@param service 当前服务名称
//@param Type 采集方式 有const, probabilistic, rateLimiting, remote 四种方式
//-- const 采样器始终对所有迹线做出相同的决定。它要么采样所有轨迹（sampler.param=1），要么不采样（sampler.param=0）。
//-- probabilistic 采样器做出随机采样决策，采样概率等于sampler.param属性值。例如，sampler.param=0.1大约有十分之一的轨迹将被采样
//-- rateLimiting 采样器使用漏斗速率限制器来确保以一定的恒定速率对轨迹进行采样。例如，sampler.param=2.0时，以每秒2条跟踪的速率对请求进行采样
//-- remote 采样器请咨询Jaeger代理以获取在当前服务中使用的适当采样策略。这允许从Jaeger后端的中央配置甚至动态地控制服务中的采样策略
//@param Param 配合Type使用的 采样概率值
//@param agentHost 代理地址:端口
func NewJaeger(service string, Type string, Param float64, agentHost string) (opentracing.Tracer, io.Closer, error) {
	if len(Type) == 0 {
		Type = "probabilistic" // 默认类型：采样器做出随机采样决策
		Param = 0.1
	}
	//github.com/uber/jaeger-client-go
	//configures and creates Jaeger Tracer 配置和创建
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{ // 采样器
			Type:  Type,
			Param: Param,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: agentHost,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, nil, err
	}
	//设置为全局的
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, err
}

//开启链路追踪
func OpenJaeger() (opentracing.Tracer, io.Closer, error) {
	isopen, err := upgin.AppConfig.Bool("jaeger.open")
	if isopen == false {
		return nil, nil, err
	}
	serviceName := upgin.AppConfig.String("jaeger.serviceName")
	if len(serviceName) == 0 {
		return nil, nil, errors.New("jaeger.serviceName is null")
	}
	jtype := upgin.AppConfig.String("jaeger.jtype")
	param, err := upgin.AppConfig.Float("jaeger.param")
	if err != nil {
		return nil, nil, err
	}
	agentHost := upgin.AppConfig.String("jaeger.agentHost")
	if len(agentHost) == 0 {
		return nil, nil, errors.New("jaeger.agentHost is empty")
	}
	return NewJaeger(serviceName, jtype, param, agentHost)
}

//中间件
//HandlerFunc defines the handler used by gin middleware as return value
func SpanMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span, cont, err := RpcxSpanWithContext(ctx.Request.Context(), fmt.Sprintf("请求地址：%s", ctx.Request.URL.Path), ctx.Request)
		if err == nil {
			defer func() {
				err2 := recover() // ****?
				if err2 != nil {
					span.SetTag("error", true)
					span.SetTag("错误信息", fmt.Sprint(err2))
				}
				span.Finish()
				if err2 != nil {
					panic(err2)
				}

			}()

			span.SetTag("请求方式", ctx.Request.Method)
			span.SetTag("get 参数", ctx.Request.URL.Query())
			method := ctx.Request.Method
			if method == "PUT" {
				raw, _ := ctx.GetRawData()
				ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
				span.SetTag("raw 参数", raw)
			} else {
				ctx.MultipartForm()
				span.SetTag("post 参数", ctx.Request.PostForm)
			}
			ctx.Request = ctx.Request.WithContext(cont)
		}
		ctx.Next()
	}

}
