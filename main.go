package main

import (
	_ "gf-simple/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "gf-simple/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-simple/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
