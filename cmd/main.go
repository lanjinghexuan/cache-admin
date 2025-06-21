package main

import (
	_ "cache-admin/config"
	"cache-admin/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	dir, _ := os.Getwd()
	fmt.Println("当前运行目录：", dir)

	r := gin.Default()
	router.Router(r)
	r.Run(":8080")

}
