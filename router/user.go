package router

import (
	"gin1/middlewares"
	"gin1/models"
	"gin1/service"

	"github.com/gin-gonic/gin"
)

// registerUserRoutes 注册用户相关路由
func registerUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/", getUser)
		userRoutes.POST("/", middlewares.ValidateRequest[models.UserRequest](), service.HandleCreateUser)
	}
}

// getUser 处理 GET 请求
func getUser(c *gin.Context) {
	service.HandleGetUser(c)
}
