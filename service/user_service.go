package service

import (
	"github.com/gin-gonic/gin"
)

// HandleGetUser 处理 GET 请求
func HandleGetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User endpoint",
	})
}

// HandleCreateUser 处理 POST 请求
func HandleCreateUser(c *gin.Context) {
	var json struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		c.JSON(200, gin.H{
			"message": "User created",
			"name":    json.Name,
			"email":   json.Email,
		})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
