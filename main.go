package main

import (
	_ "github.com/TOMO-CAT/UserManagerSystem/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/TOMO-CAT/UserManagerSystem/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
