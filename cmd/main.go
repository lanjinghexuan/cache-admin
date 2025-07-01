package main

import (
	"cache-admin/config"
	_ "cache-admin/config"
	"cache-admin/pkg"
	"cache-admin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	pkg.Init(config.Logx)
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
}
