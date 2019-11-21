package shared

import (
	"context"

	proto "github.com/damonchen/go-plugin-example/echo-example/service"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Handshake 通用的 handshake
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello-example",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"echo_grpc": &EchoPlugin{},
}

// Echo echo service
type Echo interface {
	Echo(value string) ([]byte, error)
}

// EchoPlugin echo plugin
type EchoPlugin struct {
	plugin.Plugin
	Impl Echo
}

func (p *EchoPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterEchoServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *EchoPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewEchoClient(c),
	}, nil
}
