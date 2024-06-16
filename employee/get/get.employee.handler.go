package get

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

func (handler *Handler) GetAll(c *gin.Context) {

	result, err := handler.service.GetAll()

	if err != nil {
		common.BadRequest(c, http.StatusBadRequest, "Error in getting all employees",
			[]data.ErrorDetail{
				*err,
			})
		return
	}

	common.Ok(c, http.StatusOK, "successfully got list of employees", result)
}

func (handler *Handler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BadRequest(c, http.StatusBadRequest, "Request has invalid employee id",
			[]data.ErrorDetail{
				{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: "Request has invalid employee id",
				},
			})
		return
	}
	result, errorResponse := handler.service.GetById(id)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in getting employee by id %d", id),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, "successfully got list of employees", result)
}