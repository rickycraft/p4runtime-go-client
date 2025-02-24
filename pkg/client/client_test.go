package client

import (
	"context"

	"google.golang.org/grpc"

	p4_config_v1 "github.com/p4lang/p4runtime/go/p4/config/v1"
	p4_v1 "github.com/p4lang/p4runtime/go/p4/v1"
)

type fakeP4RuntimeClient struct {
	writeFn                       func(ctx context.Context, in *p4_v1.WriteRequest, opts ...grpc.CallOption) (*p4_v1.WriteResponse, error)
	readFn                        func(ctx context.Context, in *p4_v1.ReadRequest, opts ...grpc.CallOption) (p4_v1.P4Runtime_ReadClient, error)
	setForwardingPipelineConfigFn func(ctx context.Context, in *p4_v1.SetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*p4_v1.SetForwardingPipelineConfigResponse, error)
	getForwardingPipelineConfigFn func(ctx context.Context, in *p4_v1.GetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*p4_v1.GetForwardingPipelineConfigResponse, error)
	streamChannelFn               func(ctx context.Context, opts ...grpc.CallOption) (p4_v1.P4Runtime_StreamChannelClient, error)
	capabilitiesFn                func(ctx context.Context, in *p4_v1.CapabilitiesRequest, opts ...grpc.CallOption) (*p4_v1.CapabilitiesResponse, error)
}

// fakeP4RuntimeClient implements the p4_v1.P4RuntimeClient interface
var _ p4_v1.P4RuntimeClient = &fakeP4RuntimeClient{}

func (c *fakeP4RuntimeClient) Write(ctx context.Context, in *p4_v1.WriteRequest, opts ...grpc.CallOption) (*p4_v1.WriteResponse, error) {
	if c.writeFn == nil {
		panic("No mock defined for Write RPC")
	}
	return c.writeFn(ctx, in, opts...)
}

func (c *fakeP4RuntimeClient) Read(ctx context.Context, in *p4_v1.ReadRequest, opts ...grpc.CallOption) (p4_v1.P4Runtime_ReadClient, error) {
	if c.readFn == nil {
		panic("No mock defined for Read RPC")
	}
	return c.readFn(ctx, in, opts...)
}

func (c *fakeP4RuntimeClient) SetForwardingPipelineConfig(ctx context.Context, in *p4_v1.SetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*p4_v1.SetForwardingPipelineConfigResponse, error) {
	if c.setForwardingPipelineConfigFn == nil {
		panic("No mock defined for SetForwardingPipelineConfig RPC")
	}
	return c.setForwardingPipelineConfigFn(ctx, in, opts...)
}

func (c *fakeP4RuntimeClient) GetForwardingPipelineConfig(ctx context.Context, in *p4_v1.GetForwardingPipelineConfigRequest, opts ...grpc.CallOption) (*p4_v1.GetForwardingPipelineConfigResponse, error) {
	if c.getForwardingPipelineConfigFn == nil {
		panic("No mock defined for GetForwardingPipelineConfig RPC")
	}
	return c.getForwardingPipelineConfigFn(ctx, in, opts...)
}

func (c *fakeP4RuntimeClient) StreamChannel(ctx context.Context, opts ...grpc.CallOption) (p4_v1.P4Runtime_StreamChannelClient, error) {
	if c.streamChannelFn == nil {
		panic("No mock defined for StreamChannel")
	}
	return c.streamChannelFn(ctx, opts...)
}

func (c *fakeP4RuntimeClient) Capabilities(ctx context.Context, in *p4_v1.CapabilitiesRequest, opts ...grpc.CallOption) (*p4_v1.CapabilitiesResponse, error) {
	if c.capabilitiesFn == nil {
		panic("No mock defined for Capabilities RPC")
	}
	return c.capabilitiesFn(ctx, in, opts...)
}

type fakeP4RuntimeReadClient struct {
	grpc.ClientStream
	recvFn func() (*p4_v1.ReadResponse, error)
}

// fakeP4RuntimeReadClient implements the p4_v1.P4Runtime_ReadClient interface
var _ p4_v1.P4Runtime_ReadClient = &fakeP4RuntimeReadClient{}

func (c *fakeP4RuntimeReadClient) Recv() (*p4_v1.ReadResponse, error) {
	if c.recvFn == nil {
		panic("No mock provided for Recv function")
	}
	return c.recvFn()
}

func newTestClient(p4RuntimeClient *fakeP4RuntimeClient, p4Info *p4_config_v1.P4Info) *Client {
	return &Client{
		ClientOptions:   defaultClientOptions,
		P4RuntimeClient: p4RuntimeClient,
		deviceID:        1,
		electionID:      p4_v1.Uint128{High: 0, Low: 1},
		streamSendCh:    make(chan *p4_v1.StreamMessageRequest, 1000),
		p4Info:          p4Info,
	}
}
