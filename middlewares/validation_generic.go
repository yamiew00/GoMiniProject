package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// validator
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateRequest 是一個通用middleware，用來驗證T型別的struct
func ValidateRequest[T any]() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req T

		// resolve request body
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			context.Abort()
			return
		}

		// validate model
		if err := validate.Struct(&req); err != nil {
			var validationErrors []string
			for _, err := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, err.StructNamespace()+": "+err.Tag())
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
			context.Abort()
			return
		}

		// 將驗證通過的request model 存在context中
		context.Set("validatedRequest", req)
		context.Next()
	}
}
