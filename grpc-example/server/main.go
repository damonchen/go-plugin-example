package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/damonchen/go-plugin-example/grpc-example/service"
	"google.golang.org/grpc"
)

type echoServer struct {
	pb.UnimplementedEchoServer
}

func (echo *echoServer) Echo(ctx context.Context, req *pb.EchoMesssage_Req) (*pb.EchoMesssage_Resp, error) {
	fmt.Println(string(req.Data))

	resp := pb.EchoMesssage_Resp{
		Data: req.Data,
	}

	return &resp, nil
}

func newServer() *echoServer {
	return &echoServer{}
}

func main() {
	var port string
	flag.StringVar(&port, "port", ":6789", "grpc port")
	flag.Parse()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterEchoServer(grpcServer, newServer())
	grpcServer.Serve(lis)

}
