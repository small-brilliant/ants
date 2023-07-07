package main

import (
	_ "WebServer/internal/packed"
	_ "WebServer/manifest/config"

	_ "WebServer/internal/logic"

	"WebServer/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
