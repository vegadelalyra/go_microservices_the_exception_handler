package routers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/handlers"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/models"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/token"
)

type Login struct {
	Logger *log.Logger
	loginHandler *handlers.Login
	flags *models.Flags
}

func NewRoute(logger *log.Logger, flags *models.Flags) *Login {
	loginHandler := handlers.NewLogin(logger, flags)
	token.Init()

	return &Login{
		Logger: logger,
		loginHandler: loginHandler,
		flags: flags,
	}
}

func (r *Login) RegisterRoutes() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.POST("/login", r.loginHandler.Login)
	ginEngine.POST("/verifyToken", r.loginHandler.VerifyToken)
	return ginEngine 
}