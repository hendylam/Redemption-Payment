package v1

import (
	"context"

	pool "github.com/processout/grpc-go-pool"
)

type (
	// GrpcClientManager interface
	GrpcClientManager interface {
		Client(ctx context.Context) (GrpcMitraintegrasiClient, error)
	}

	grpcClientManager struct {
		connPool *pool.Pool
	}

	// GrpcMitraintegrasiClient struct
	GrpcMitraintegrasiClient interface {
		Close()
		SendRedemPayDataClient
	}

	grpcMitraintegrasiClient struct {
		conn *pool.ClientConn
		*sendRedemPayDataClient
	}
)

// Close
func (c *grpcMitraintegrasiClient) Close() {
	c.conn.Close()
}

// NewGrpcClientManager func
func NewGrpcClientManager(connPool *pool.Pool) GrpcClientManager {
	return &grpcClientManager{
		connPool: connPool,
	}
}

// Client func
func (c *grpcClientManager) Client(ctx context.Context) (GrpcMitraintegrasiClient, error) {
	cn, err := c.connPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	client := &grpcMitraintegrasiClient{
		conn:                   cn,
		sendRedemPayDataClient: NewSendRedemPayDataClient(cn).(*sendRedemPayDataClient),
	}
	return client, nil
}
