package main

import (
	"fmt"
	"github.com/damonchen/go-plugin-example/echo-example/shared"
	"github.com/hashicorp/go-plugin"
	"os"
	"os/exec"
)

func main() {

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		// 注意，此处的 ECHO_PLUGIN 是外部指定，实际上，这儿可以写一个函数，用来处理各种路径
		Cmd: exec.Command("sh", "-c", os.Getenv("ECHO_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
	})
	defer client.Kill()

	grpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// 这儿 echo_grpc， 是前面 PluginMap 中的 key
	raw, err := grpcClient.Dispense("echo_grpc")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a Echo now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	echo := raw.(shared.Echo)
	args := os.Args[1]
	bytes, err := echo.Echo(args)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("echo: '%s'\n", string(bytes))

	os.Exit(0)
}
