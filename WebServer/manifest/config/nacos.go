package config

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/contrib/config/nacos/v2"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Nacos struct {
	ServerConfig []constant.ServerConfig
	ClientConfig constant.ClientConfig
	ConfigParam  vo.ConfigParam
}

func init() {
	ctx := gctx.GetInitCtx()
	nacosCfg := &Nacos{}
	v, err := g.Cfg().Get(ctx, "nacos")
	if err != nil {
		fmt.Println("err:", err)
	}
	json.Unmarshal(v.Bytes(), &nacosCfg)
	// Create anacosClient that implements gcfg.Adapter.
	adapter, err := nacos.New(ctx, nacos.Config{
		ServerConfigs: nacosCfg.ServerConfig,
		ClientConfig:  nacosCfg.ClientConfig,
		ConfigParam:   nacosCfg.ConfigParam,
	})
	if err != nil {
		g.Log().Fatalf(ctx, `%+v`, err)
	}
	// Change the adapter of default configuration instance.
	g.Cfg().SetAdapter(adapter)
}
