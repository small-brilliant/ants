package cmd

import (
	"WebServer/internal/controller"
	"context"
	"encoding/json"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			d := g.Cfg().MustGet(ctx, "etcd")
			clientCfg := clientv3.Config{}
			json.Unmarshal(d.Bytes(), &clientCfg)
			c, _ := clientv3.New(clientCfg)
			grpcx.Resolver.Register(etcd.NewWithClient(c))
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Auth,
				)
			})
			s.Run()
			return nil
		},
	}
)
