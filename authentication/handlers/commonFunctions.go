package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/models"
)


func ok(context *gin.Context, status int, message string, data interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, models.Response{
		Data:    data,
		Status:  status,
		Message: message,
	})
}

func badRequest(context *gin.Context, status int, message string, errors []models.ErrorDetail) {
	context.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
		Error:   errors,
		Status:  status,
		Message: message,
	})
}

func returnUnauthorized(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error: []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeUnauthorized,
				ErrorMessage: "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}