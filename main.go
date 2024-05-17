package main

import (
	"gin1/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 Gin router
	r := gin.Default()

	// 注册路由
	router.RegisterRoutes(r)

	// 启动 Gin 服务器，默认监听在 0.0.0.0:8080
	r.Run() // listen and serve on 0.0.0.0:8080
}
