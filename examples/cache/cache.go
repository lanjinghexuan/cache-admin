package cache

import (
	_ "cache-admin/config"
	"cache-admin/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	Name string
	Age  int
}

func GetUser(page int32, limit int32) ([]*User, error) {
	fmt.Println(page, limit)
	var user []*User
	user = append(user, &User{
		Name: "张三",
		Age:  18,
	})
	user = append(user, &User{
		Name: "李四",
		Age:  20,
	})
	return user, nil
}

type GetUserReq struct {
	Page  int32 `json:"page" form:"page"`
	Limit int32 `json:"limit" form:"limit"`
}

func GetUserCache(c *gin.Context) {
	var req GetUserReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	params := map[string]interface{}{"page": req.Page, "limit": req.Limit}
	var res interface{}
	err = pkg.GetCache(&res, pkg.CacheData{
		Prefix: "cache-admin/GetUserCache",
		Params: params,
		Exprie: 86400 * time.Second,
	}, func() (interface{}, error) {
		return GetUser(req.Page, req.Limit)
	})
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
	return
}
