package auth_grpc_client

import "github.com/stretchr/testify/mock"

type AuthGrpcClientMock struct {
	mock.Mock
}

func (a *AuthGrpcClientMock) TerminateProfile(email string) error {
	args := a.Called(email)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}
