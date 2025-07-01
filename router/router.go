package router

import (
	"cache-admin/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	c := r.Group("/cache")
	{
		c.GET("/del", handler.CacheDel)
		c.GET("/delByPrefix", handler.CacheDelByPrefix)
		c.GET("/getCacheKeyList/:prefix", handler.GetCacheKeyList)
		c.GET("/cacheDelFind", handler.CacheDelFind)
	}
}
