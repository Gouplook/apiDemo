/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/11 13:46
@Description:

*********************************************/
package client

import (
	"apiDemo/upbase/common/plugins/jaeger"
	"apiDemo/upgin"
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"net/http"
	"sync"
)

var (
	rpcPools map[string]map[string]*client.XClientPool
	lock     *sync.RWMutex
)

type Baseclient struct {
	ServiceName string
	ServicePath string
	discovery   client.ServiceDiscovery // rpcx 服务发现
	xClient     client.XClient          // 用于客户端的服务发现与治理
}

func init() {
	rpcPools = map[string]map[string]*client.XClientPool{}
	lock = new(sync.RWMutex)
}

func (cli *Baseclient) getPools(serviceName string, servicePath string) client.XClient {
	if service, ok := rpcPools[serviceName]; ok {
		if rpcpool, ok := service[servicePath]; ok {
			return rpcpool.Get()
		} else {
			lock.Lock()
			rpcpool, ok := service[servicePath]
			if !ok {
				rpcpool = client.NewXClientPool(upgin.AppConfig.DefaultInt("rpc_pool_count", 10), cli.ServicePath, client.Failtry, client.RandomSelect, cli.GetDiscovery(), client.DefaultOption)
				rpcPools[serviceName][servicePath] = rpcpool
			}
			lock.Unlock()
			return rpcpool.Get()
		}
	} else {
		lock.Lock()
		service, ok := rpcPools[serviceName]
		if !ok {
			service = map[string]*client.XClientPool{
				servicePath: client.NewXClientPool(upgin.AppConfig.DefaultInt("rpc_pool_count", 10), cli.ServicePath, client.Failtry, client.RandomSelect, cli.GetDiscovery(), client.DefaultOption),
			}
			rpcPools[serviceName] = service
		}
		lock.Unlock()
		return service[servicePath].Get()
	}
}

// 获取服务发现
func (cli *Baseclient) GetDiscovery() client.ServiceDiscovery {
	if cli.discovery == nil {
		address := upgin.AppConfig.String(cli.ServiceName)
		cli.discovery, _ = client.NewPeer2PeerDiscovery(address, "")
	}
	return cli.discovery
}
func (cli *Baseclient) getXClient() client.XClient {
	return cli.getPools(cli.ServiceName, cli.ServicePath)
}

// 请求
func (cli *Baseclient) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	// 操作方法的Name
	operationName := fmt.Sprintf("调用%s服务的%s方法", cli.ServicePath, serviceMethod)
	span, ctx, spanErr := jaeger.RpcxSpanWithContext(ctx, operationName, &http.Request{})
	if spanErr == nil {
		span.SetTag("参数", args)
		defer span.Finish()
	}

	// 调用  XClient interface 中call 接口
	err := cli.getXClient().Call(ctx, serviceMethod, args, reply)
	if err != nil && spanErr == nil {
		span.SetTag("error", true)
		span.SetTag("错误信息", fmt.Sprint(err))
	}
	return err

}
func (cli *Baseclient)Close() {

}
