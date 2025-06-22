package main

import (
	_ "cache-admin/config"
	"cache-admin/examples/cache"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/getUser", cache.GetUserCache)
	r.Run(":8081")
}
