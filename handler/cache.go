package handler

import (
	"cache-admin/config"
	"cache-admin/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CacheReq struct {
	Prefix string                 `json:"prefix" form:"prefix"`
	Params map[string]interface{} `json:"params" form:"params"`
}

func CacheDel(c *gin.Context) {
	var req CacheReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cachedata pkg.CacheData
	cachedata.Prefix = req.Prefix
	cachedata.Params = req.Params
	err = pkg.CacheDel(cachedata)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": "删除成功",
	})
}

type CacheDelByPrefixReq struct {
	Prefix string `json:"prefix" form:"prefix" binding:"required"`
}

func CacheDelByPrefix(c *gin.Context) {
	var req CacheDelByPrefixReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	err = pkg.CacheDelByPrefix(req.Prefix)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
	return
}

type CacheDelByPrefixResp struct {
	//Prefix string `json:"prefix" form:"prefix"`
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func GetCacheKeyList(c *gin.Context) {
	var req CacheDelByPrefixResp
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Limit == 0 {
		req.Limit = 100
	}
	if req.Page == 0 {
		req.Page = 1
	}
	prefix := c.Param("prefix")
	iter := config.RedisDB.Scan(config.Ctx, 0, prefix+"*", 0).Iterator()
	var keys []string
	for iter.Next(config.Ctx) {
		keys = append(keys, iter.Val())
	}
	start := (req.Page - 1) * req.Limit
	end := start + req.Limit
	total := len(keys)
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pageKeys := keys[start:end]

	var results []map[string]interface{}
	for _, key := range pageKeys {
		ttl, _ := config.RedisDB.TTL(config.Ctx, key).Result()
		size, _ := config.RedisDB.MemoryUsage(config.Ctx, key).Result()
		results = append(results, map[string]interface{}{
			"key":  key,
			"ttl":  int(ttl.Seconds()),
			"size": size,
		})
	}

	c.JSON(200, gin.H{
		"data":  results,
		"total": total,
	})
}

type CacheDelFindReq struct {
	Prefix string `json:"prefix" form:"prefix"`
}

func CacheDelFind(c *gin.Context) {
	var req CacheDelFindReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.RedisDB.Del(config.Ctx, req.Prefix)

	c.JSON(200, gin.H{
		"code": 200,
		"data": "删除成功",
	})
}
