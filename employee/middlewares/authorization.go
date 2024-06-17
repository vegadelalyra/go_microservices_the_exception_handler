package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/employee/data"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/rpc/rpc_auth"
	"google.golang.org/grpc"
)

var validateClient rpc_auth.ValidateTokenServiceClient

func InitMiddleware() *data.ErrorDetail {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		return &data.ErrorDetail{
			ErrorMessage: fmt.Sprintf("error in dial %+v", err),
			ErrorType:    data.ErrorTypeFatal,
		}
	}
	validateClient = rpc_auth.NewValidateTokenServiceClient(conn)
	return nil
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(gContext *gin.Context) {
		tokenString := gContext.Request.Header.Get("apikey")

		if tokenString == "" {
			ReturnUnauthorized(gContext)
		}

		fmt.Printf("token --%s--", tokenString)
		response, err := validateClient.SimpleValidate(context.Background(), &rpc_auth.ValidateTokenRequest{
			Token: tokenString,
		})

		if err != nil {
			ReturnUnauthorized(gContext)
			return
		}
		if len(gContext.Keys) == 0 {
			gContext.Keys = make(map[string]interface{})
		}
		gContext.Keys["CompanyId"] = response.CompanyId
		gContext.Keys["Username"] = response.Username
		gContext.Keys["Roles"] = response.Roles
	}
}

func ReturnUnauthorized(context *gin.Context) {}
