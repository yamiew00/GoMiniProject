// service/user_service.go
package service

import (
	"gin1/models"

	"github.com/gin-gonic/gin"
)

// HandleGetUser 处理 GET 请求
func HandleGetUser(c *gin.Context) {
	c.JSON(200, models.UserResponse{
		Message: "User endpoint",
	})
}

// HandleCreateUser 处理 POST 请求并从上下文中获取请求体
func HandleCreateUser(c *gin.Context) {
	// 从上下文中获取验证通过的请求体
	validatedRequest, exists := c.Get("validatedRequest")
	if !exists {
		c.JSON(500, gin.H{"error": "Validation middleware not working properly"})
		return
	}

	req := validatedRequest.(models.UserRequest)

	c.JSON(200, models.UserResponse{
		Message: "User created",
		Name:    req.Name,
		Email:   req.Email,
	})
}
