package handler

import (
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
	Prefix string `json:"prefix" form:"prefix"`
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
