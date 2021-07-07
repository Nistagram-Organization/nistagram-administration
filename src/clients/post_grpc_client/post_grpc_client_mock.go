package post_grpc_client

import (
	"github.com/stretchr/testify/mock"
)

type PostGrpcClientMock struct {
	mock.Mock
}

func (c *PostGrpcClientMock) DecideOnPost(id uint, delete bool) error {
	args := c.Called(id, delete)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}