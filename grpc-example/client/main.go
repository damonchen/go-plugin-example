package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/damonchen/go-plugin-example/grpc-example/service"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:6789", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := pb.EchoMesssage_Req{
		Data: []byte("hello world"),
	}

	resp, err := client.Echo(ctx, &req)
	if err != nil {
		log.Fatalf("fail to echo: %v", err)
	}

	fmt.Printf("%s\n", resp.Data)

}
