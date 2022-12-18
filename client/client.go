package client

import (
	"context"
	"fmt"
	"github.com/fullstackwang/tron-grpc/wallet"

	"github.com/fullstackwang/tron-grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

//go:generate go run ../tools/generator/gen.go

type Client struct {
	Signer  Signer
	address string
	conn    *grpc.ClientConn
	client  api.WalletClient
	timeout time.Duration
	opts    []grpc.DialOption
	apiKey  string
}

func New(address, apikey string) *Client {
	return NewWithTimeout(address, apikey, 5*time.Second)
}

func NewWithTimeout(address, apikey string, timeout time.Duration) *Client {
	client := &Client{
		address: address,
		timeout: timeout,
		apiKey:  apikey,
	}
	return client
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) SetPrivateKey(key string) {
	c.Signer = wallet.NewWalletFromPrivateKey(key)
}

// SetTimeout for Client connections
func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// Start initiate grpc  connection
func (c *Client) Start(opts ...grpc.DialOption) error {
	var err error
	if c.address == "" {
		c.address = "grpc.trongrid.io:50051"
	}
	c.opts = opts
	c.conn, err = grpc.Dial(c.address, opts...)

	if err != nil {
		return fmt.Errorf("Connecting GRPC Client: %v", err)
	}
	c.client = api.NewWalletClient(c.conn)
	return nil
}

// Stop GRPC Connection
func (c *Client) Stop() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// Reconnect GRPC
func (c *Client) Reconnect(url string) error {
	c.Stop()
	if url != "" {
		c.address = url
	}
	return c.Start(c.opts...)
}

func (c *Client) makeContext(parent context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, c.timeout)
	if c.apiKey != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "TRON-PRO-API-KEY", c.apiKey)
	}
	return ctx, cancel
}
