package router

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// 注册用户路由
	registerUserRoutes(r)
	// 注册产品路由
	registerProductRoutes(r)
}
