package transport

import (
	"context"
	"crm_system/config/crm_core"
	pb "crm_system/pkg/auth/authservice/gw"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ValidateGrpcTransport struct {
	config crm_core.Configuration
	client pb.AuthServiceClient
}

func NewValidateGrpcTransport(config crm_core.Configuration) *ValidateGrpcTransport {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, _ := grpc.Dial(config.Transport.ValidateGrpc.Host, opts...)

	client := pb.NewAuthServiceClient(conn)

	return &ValidateGrpcTransport{
		client: client,
		config: config,
	}
}

func (t *ValidateGrpcTransport) ValidateTransport(ctx context.Context, accessToken string, roles ...string) (*pb.ResponseJSON, error) {

	resp, err := t.client.Validate(ctx, &pb.ValidateRequest{
		AccessToken: accessToken,
		Roles:       roles,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot Validate: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("not found")
	}

	return resp.Response, nil
}
