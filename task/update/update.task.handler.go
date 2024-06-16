package update

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/task/common"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/task/data"
)

type Handler struct {
	service *Service
}

func InitHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Update(c *gin.Context) {
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
	var updateObj data.Task
	if err := c.ShouldBindJSON(&updateObj); err != nil {
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
	updateObj.ID = id

	result, errorResponse := h.service.Update(&updateObj)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in updating task by id %d", id),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully updated task with id: %d", id), result)
}