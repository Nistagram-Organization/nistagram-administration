package administration

import (
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/auth_grpc_client"
	"github.com/Nistagram-Organization/nistagram-shared/src/utils/rest_error"
)

type AdministrationService interface {
	TerminateProfile(email string) rest_error.RestErr
}

type administrationService struct {
	auth_grpc_client.AuthGrpcClient
}

func NewAdministrationService(authGrpcClient auth_grpc_client.AuthGrpcClient) AdministrationService {
	return &administrationService{
		authGrpcClient,
	}
}

func (s *administrationService) TerminateProfile(email string) rest_error.RestErr {
	if err := s.AuthGrpcClient.TerminateProfile(email); err != nil {
		return rest_error.NewInternalServerError("auth grpc client error when terminating profile", err)
	}
	return nil
}
