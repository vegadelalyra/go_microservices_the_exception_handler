package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/task/data"
)

func Ok(context *gin.Context, status int, message string, result interface{}) {
	context.AbortWithStatusJSON(status,  data.Response{
		Data:    result,
		Status:  status,
		Message: message,
	})
}
func BadRequest(context *gin.Context, status int, message string, errors []data.ErrorDetail) {
	context.AbortWithStatusJSON(status, data.Response{
		Error:   errors,
		Status:  status,
		Message: message,
	})
}

func ReturnUnauthorized(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, data.Response{
		Error: []data.ErrorDetail{
			{
				ErrorType:    data.ErrorTypeUnauthorized,
				ErrorMessage: "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}