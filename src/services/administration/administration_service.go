package administration

import (
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/auth_grpc_client"
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/post_grpc_client"
	"github.com/Nistagram-Organization/nistagram-administration/src/dtos/inappropriate_post_report_decision"
	"github.com/Nistagram-Organization/nistagram-shared/src/utils/rest_error"
)

type AdministrationService interface {
	DecideOnPost(*inappropriate_post_report_decision.InappropriatePostReportDecision) (uint, rest_error.RestErr)
}

type administrationService struct {
	auth_grpc_client.AuthGrpcClient
	post_grpc_client.PostGrpcClient
}

func NewAdministrationService(authGrpcClient auth_grpc_client.AuthGrpcClient, postGrpcClient post_grpc_client.PostGrpcClient) AdministrationService {
	return &administrationService{
		authGrpcClient,
		postGrpcClient,
	}
}

func (s *administrationService) DecideOnPost(decision *inappropriate_post_report_decision.InappropriatePostReportDecision) (uint, rest_error.RestErr) {
	if err := decision.Validate(); err != nil {
		return 0, err
	}

	if err := s.PostGrpcClient.DecideOnPost(decision.PostID, decision.Delete); err != nil {
		return 0, rest_error.NewInternalServerError("Failed to decide on post", err)
	}

	if decision.Terminate {
		if err := s.AuthGrpcClient.TerminateProfile(decision.AuthorEmail); err != nil {
			return 0, rest_error.NewInternalServerError("Failed to terminate user profile", err)
		}
	}

	return decision.PostID, nil
}
