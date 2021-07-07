package administration

import (
	"errors"
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/auth_grpc_client"
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/post_grpc_client"
	"github.com/Nistagram-Organization/nistagram-administration/src/dtos/inappropriate_post_report_decision"
	"github.com/Nistagram-Organization/nistagram-shared/src/utils/rest_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AdministrationServiceUnitTestsSuite struct {
	suite.Suite
	postGrpcClientMock    *post_grpc_client.PostGrpcClientMock
	authGrpcClientMock    *auth_grpc_client.AuthGrpcClientMock
	service                AdministrationService
}

func TestAdministrationServiceUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(AdministrationServiceUnitTestsSuite))
}

func (suite *AdministrationServiceUnitTestsSuite) SetupSuite() {
	suite.postGrpcClientMock = new(post_grpc_client.PostGrpcClientMock)
	suite.authGrpcClientMock = new(auth_grpc_client.AuthGrpcClientMock)
	suite.service = NewAdministrationService(suite.authGrpcClientMock, suite.postGrpcClientMock)
}

func (suite *AdministrationServiceUnitTestsSuite) TestNewAdministrationService() {
	assert.NotNil(suite.T(), suite.service, "Service is nil")
}

func (suite *AdministrationServiceUnitTestsSuite) TestAdministrationService_DecideOnPost_CannotTerminateAuthor() {
	decision := inappropriate_post_report_decision.InappropriatePostReportDecision{
		PostID:      1,
		AuthorEmail: "",
		Delete:      false,
		Terminate:   true,
	}
	err := rest_error.NewBadRequestError("Cannot terminate author's profile and not delete reported post")

	_, decideErr := suite.service.DecideOnPost(decision)

	assert.Equal(suite.T(), err, decideErr)
}

func (suite *AdministrationServiceUnitTestsSuite) TestAdministrationService_DecideOnPost_PostGrpcError() {
	decision := inappropriate_post_report_decision.InappropriatePostReportDecision{
		PostID:      1,
		AuthorEmail: "",
		Delete:      true,
		Terminate:   true,
	}
	errGrpc := errors.New("")
	err := rest_error.NewInternalServerError("Failed to decide on post", errGrpc)

	suite.postGrpcClientMock.On("DecideOnPost", decision.PostID, decision.Delete).Return(errGrpc).Once()

	_, decideErr := suite.service.DecideOnPost(decision)

	assert.Equal(suite.T(), err, decideErr)
}


func (suite *AdministrationServiceUnitTestsSuite) TestAdministrationService_DecideOnPost_AuthGrpcError() {
	decision := inappropriate_post_report_decision.InappropriatePostReportDecision{
		PostID:      1,
		AuthorEmail: "",
		Delete:      true,
		Terminate:   true,
	}
	errGrpc := errors.New("")
	err := rest_error.NewInternalServerError("Failed to terminate user profile", errGrpc)

	suite.postGrpcClientMock.On("DecideOnPost", decision.PostID, decision.Delete).Return(nil).Once()
	suite.authGrpcClientMock.On("TerminateProfile", decision.AuthorEmail).Return(errGrpc).Once()

	_, decideErr := suite.service.DecideOnPost(decision)

	assert.Equal(suite.T(), err, decideErr)
}

func (suite *AdministrationServiceUnitTestsSuite) TestAdministrationService_DecideOnPost() {
	decision := inappropriate_post_report_decision.InappropriatePostReportDecision{
		PostID:      1,
		AuthorEmail: "",
		Delete:      true,
		Terminate:   true,
	}

	suite.postGrpcClientMock.On("DecideOnPost", decision.PostID, decision.Delete).Return(nil).Once()
	suite.authGrpcClientMock.On("TerminateProfile", decision.AuthorEmail).Return(nil).Once()

	retVal, decideErr := suite.service.DecideOnPost(decision)

	assert.Equal(suite.T(), decision.PostID, retVal)
	assert.Equal(suite.T(), nil, decideErr)
}
