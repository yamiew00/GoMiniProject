package router

import (
	"github.com/gin-gonic/gin"
)

// registerProductRoutes 注册产品相关路由
func registerProductRoutes(r *gin.Engine) {
	productRoutes := r.Group("/product")
	{
		productRoutes.GET("/", getProduct)
	}
}

// getProduct 处理 GET 请求
func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Product endpoint",
	})
}
