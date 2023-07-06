package main

import (
	_ "AuthServer/internal/packed"

	_ "AuthServer/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"AuthServer/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
