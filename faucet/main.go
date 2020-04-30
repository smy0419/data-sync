package main

import (
	"github.com/AsimovNetwork/data-sync/faucet/handler"
	"github.com/AsimovNetwork/data-sync/faucet/router"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(common.Cfg.GinMode)
	root := gin.New()
	root.Use(handler.UseLog(), gin.Recovery())

	// 跨域
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	root.Use(cors.New(config))

	root.POST("/", router.FaucetTransfer)

	root.Run(common.Cfg.Port)
}
