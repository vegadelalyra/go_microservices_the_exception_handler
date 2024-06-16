package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/models"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/services"
)

type Login struct {
	logger *log.Logger
	flags *models.Flags 
	loginService *services.Login
}

func NewLogin(logger *log.Logger, flags *models.Flags) *Login {
	loginService := services.NewLogin(logger, flags)
	return &Login{
		logger: logger, 
		flags: flags,
		loginService: loginService,
	}
}

func (l *Login) Login(context *gin.Context) {

	var loginObj models.LoginRequest
	if err := context.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType: models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
		return
	}

	tokenString, err := l.loginService.GetToken(loginObj, context.Request.Header.Get("Referer"))
	if err != nil {
		badRequest(context, http.StatusBadRequest, "error in generating token", []models.ErrorDetail{ *err, })
		return
	}
	
	ok(context, http.StatusOK, "token created", tokenString)
}

func (l *Login) VerifyToken(context *gin.Context) {
	tokenString := context.Request.Header.Get("apikey")
	referer := context.Request.Header.Get("Referer")

	if tokenString == "" {
		returnUnauthorized(context)
		return
	}	

	valid, claims := l.loginService.VerifyToken(tokenString, referer)
	if !valid {
		returnUnauthorized(context)
		return
	}
	ok(context, http.StatusOK, "token is valid", claims)
}