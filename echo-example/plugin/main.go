package main

import (
	"github.com/damonchen/go-plugin-example/echo-example/shared"
	"github.com/hashicorp/go-plugin"
)

// 实现 echo 的服务
type Echo struct{}

func (Echo) Echo(value string) ([]byte, error) {
	return []byte(value), nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			// 注意，此处的名称，应该和 client 端的 echo_grpc的名称保持一致
			"echo_grpc": &shared.EchoPlugin{Impl: &Echo{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
