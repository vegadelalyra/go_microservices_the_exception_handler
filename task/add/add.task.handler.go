package add

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/task/common"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/task/data"
)

type Handler struct{
	service *Service
}

func InitHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}


func (handler * Handler) Add(c *gin.Context){
	var addObj data.Task
	if err := c.ShouldBindJSON(&addObj); err != nil {
		common.BadRequest(c, http.StatusBadRequest, "Request has invalid body",
			[]data.ErrorDetail{
				{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: "Request has invalid body",
				},
				{
					ErrorType:    data.ErrorTypeValidation,
					ErrorMessage: err.Error(),
				},
			})
		return
	}
	result, errorResponse := handler.service.Add(&addObj)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in Adding task by name %s", addObj.Name),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully Added task with name %s", addObj.Name), result)
}