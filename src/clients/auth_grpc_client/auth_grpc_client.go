package auth_grpc_client

import (
	"context"
	"github.com/Nistagram-Organization/nistagram-shared/src/proto"
	"google.golang.org/grpc"
)

type AuthGrpcClient interface {
	TerminateProfile(email string) error
}

type authGrpcClient struct {
}

func NewAuthGrpcClient() AuthGrpcClient {
	return &authGrpcClient{}
}

func (c *authGrpcClient) TerminateProfile(email string) error {
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := proto.NewAuthServiceClient(conn)

	r, err := client.TerminateProfile(ctx,
		&proto.TerminateProfileRequest{
			Email: email,
		},
	)

	if err != nil || !r.Success {
		return err
	}

	return nil
}
