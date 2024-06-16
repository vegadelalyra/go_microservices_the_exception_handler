package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/models"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/repository"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/token"
)

type Login struct {
	logger          *log.Logger
	flags           *models.Flags
	loginRepository *repository.Login
}

func NewLogin(logger *log.Logger, flags *models.Flags) *Login {
	return &Login{
		logger:          logger,
		flags:           flags,
		loginRepository: repository.Init(),
	}
}

func (l *Login) GetToken(loginModel models.LoginRequest, origin string) (string, *models.ErrorDetail) {
	user, err := l.loginRepository.GetUserByUserName(loginModel.UserName, loginModel.Password)
	if err != nil {
		return "", &models.ErrorDetail{
			ErrorType: models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		}
	}
	var claims = &models.JwtClaims{
		CompanyId: strconv.Itoa(user.Id),
		Username: user.Name,
		Roles: user.Roles,
		StandardClaims: jwt.StandardClaims{
			Audience: origin,
		},
	}
	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	return token.GenrateToken(claims, expirationTime)

}

func (*Login) VerifyToken(tokenString, referer string) (bool, *models.JwtClaims) {
	return token.VerifyToken(tokenString, referer)
}