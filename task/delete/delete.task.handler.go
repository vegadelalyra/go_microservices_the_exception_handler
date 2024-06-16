package delete

import (
	"fmt"
	"net/http"
	"strconv"

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


func (handler * Handler) Delete(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BadRequest(c, http.StatusBadRequest, "Request has invalid task id",
			[]data.ErrorDetail{
				{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: "Request has invalid task id",
				},
			})
		return
	}
	result, errorResponse := handler.service.Delete(id)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in deleting task by id %d", id),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully deleted task with id: %d", id), result)
}