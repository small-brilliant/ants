package cmd

import (
	"AuthServer/internal/controller/user"
	"context"
	"encoding/json"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start user grpc server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			d := g.Cfg().MustGet(ctx, "etcd")
			clientCfg := clientv3.Config{}
			json.Unmarshal(d.Bytes(), &clientCfg)
			c, _ := clientv3.New(clientCfg)
			grpcx.Resolver.Register(etcd.NewWithClient(c))
			s := grpcx.Server.New()
			user.Register(s)
			s.Run()
			return nil
		},
	}
)
