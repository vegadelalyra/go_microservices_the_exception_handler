package update

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/employee/common"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/employee/data"
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
		common.BadRequest(c, http.StatusBadRequest, "request has invalid employee id",
			[]data.ErrorDetail{
				{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: "request has invalid employee id",
				},
			})
		return
	}
	var updateObj data.Employee
	if err := c.ShouldBindJSON(&updateObj); err != nil {
		common.BadRequest(c, http.StatusBadRequest, "request has invalid body",
			[]data.ErrorDetail{
				{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: "request has invalid body",
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
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in updating employee by id %d", id),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully updated employees with id: %d", id), result)
}