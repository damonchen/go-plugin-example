package shared

import (
	"context"

	proto "github.com/damonchen/go-plugin-example/echo-example/service"
)

type GRPCClient struct {
	client proto.EchoClient
}

// 实现 client 端的 echo 接口
func (c *GRPCClient) Echo(value string) ([]byte, error) {
	resp, err := c.client.Echo(context.Background(), &proto.GetRequest{
		Key: value,
	})
	return resp.Value, err
}

// GRPCServer grpc server
type GRPCServer struct {
	Impl Echo
}

// Echo impl
func (s *GRPCServer) Echo(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := s.Impl.Echo(req.Key)
	return &proto.GetResponse{
		Value: v,
	}, err
}
