package tron_grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GrpcClient controller structure
type GrpcClient struct {
	Address string
	Conn    *grpc.ClientConn
	timeout time.Duration
	opts    []grpc.DialOption
	apiKey  string
}

// NewGrpcClient create grpc controller
func NewGrpcClient(address, apikey string) *GrpcClient {
	return NewGrpcClientWithTimeout(address, apikey, 5*time.Second)
}

// NewGrpcClientWithTimeout create grpc controller
func NewGrpcClientWithTimeout(address, apikey string, timeout time.Duration) *GrpcClient {
	client := &GrpcClient{
		Address: address,
		timeout: timeout,
		apiKey:  apikey,
	}
	return client
}

// SetTimeout for Client connections
func (g *GrpcClient) SetTimeout(timeout time.Duration) {
	g.timeout = timeout
}

// Start initiate grpc  connection
func (g *GrpcClient) Start(opts ...grpc.DialOption) error {
	var err error
	if g.Address == "" {
		g.Address = "grpc.trongrid.io:50051"
	}
	g.opts = opts
	g.Conn, err = grpc.Dial(g.Address, opts...)

	if err != nil {
		return fmt.Errorf("Connecting GRPC Client: %v", err)
	}
	return nil
}

func (g *GrpcClient) MakeContext(parent context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, g.timeout)
	if g.apiKey != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "TRON-PRO-API-KEY", g.apiKey)
	}
	return ctx, cancel
}

// Stop GRPC Connection
func (g *GrpcClient) Stop() {
	if g.Conn != nil {
		g.Conn.Close()
	}
}

// Reconnect GRPC
func (g *GrpcClient) Reconnect(url string) error {
	g.Stop()
	if url != "" {
		g.Address = url
	}
	return g.Start(g.opts...)
}
