package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/subosito/drafter-go/plugin/adapter"
	rc "github.com/subosito/drafter-go/rpc-plugin/rpc"
)

func main() {
	client := adapter.New()

	var pluginMap = map[string]plugin.Plugin{
		"drafter": &rc.Plugin{Impl: client},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: rc.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
