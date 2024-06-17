package services

import (
	"io"
	"log"
	"os"

	"github.com/vegadelalyra/go_microservices_the_exception_handler/auth/models"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/rpc/rpc_auth"
)

type ValidateRpcServer struct {
	rpc_auth.UnimplementedValidateTokenServiceServer
}

func (ValidateRpcServer) Validate(stream rpc_auth.ValidateTokenService_ValidateServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		logger := log.New(os.Stdout, "validateRpc", 1)
		flags, errF := models.GetFlags()
		if errF != nil {
			return errF
		}

		service := NewLogin(logger, flags)

		valid, claims := service.VerifyToken(in.Token, "")

		if !valid {
			if err := stream.Send(&rpc_auth.ValidateTokenResponse{
				IsValid:   valid,
				CompanyId: "",
				Username:  "",
				Roles:     nil,
			}); err != nil {
				return err
			}
		} else {
			if err := stream.Send(&rpc_auth.ValidateTokenResponse{
				IsValid:   valid,
				CompanyId: claims.CompanyId,
				Username:  claims.Username,
				Roles:     nil,
			}); err != nil {
				return err
			}
		}
	}
}
