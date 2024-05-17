package router

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	registerUserRoutes(r)
	registerProductRoutes(r)
}
