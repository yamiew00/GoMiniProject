package service

import (
	"gin1/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// validator 实例
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// HandleGetUser 处理 GET 请求
func HandleGetUser(c *gin.Context) {
	c.JSON(200, models.UserResponse{
		Message: "User endpoint",
	})
}

// HandleCreateUser 处理 POST 请求并验证请求体
func HandleCreateUser(c *gin.Context) {
	var req models.UserRequest

	// 绑定 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, models.UserResponse{Message: "Invalid request"})
		return
	}

	// 验证请求体
	if err := validate.Struct(&req); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.StructNamespace()+": "+err.Tag())
		}
		c.JSON(400, gin.H{"error": validationErrors})
		return
	}

	c.JSON(200, models.UserResponse{
		Message: "User created",
		Name:    req.Name,
		Email:   req.Email,
	})
}
