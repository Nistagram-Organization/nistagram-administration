package post_grpc_client

import (
	"context"
	"github.com/Nistagram-Organization/nistagram-shared/src/proto"
	"google.golang.org/grpc"
)

type PostGrpcClient interface {
	DecideOnPost(id uint, delete bool) error
}

type postGrpcClient struct {
}

func NewPostGrpcClient() PostGrpcClient {
	return &postGrpcClient{}
}

func (c *postGrpcClient) DecideOnPost(id uint, delete bool) error {
	conn, err := grpc.Dial("127.0.0.1:8085", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := proto.NewPostServiceClient(conn)

	r, err := client.DecideOnPost(ctx,
		&proto.DecideOnPostRequest{
			Post:   uint64(id),
			Delete: delete,
		},
	)

	if err != nil || !r.Success {
		return err
	}

	return nil
}
