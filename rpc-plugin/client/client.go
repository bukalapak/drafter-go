package client

import (
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/subosito/drafter-go"
	"github.com/subosito/drafter-go/rpc-plugin/rpc"
)

type DrafterRPC interface {
	Dispense() (drafter.Drafter, error)
	Close()
}

type engine struct {
	client *plugin.Client
}

type Config struct {
	Name   string
	Logger hclog.Logger
}

func (c Config) cmdName() string {
	if c.Name == "" {
		return "drafter-rpc"
	}

	return c.Name
}

func (c Config) logger() hclog.Logger {
	if c.Logger == nil {
		return hclog.NewNullLogger()
	}

	return c.Logger
}

func New(cfg Config) DrafterRPC {
	pluginMap := map[string]plugin.Plugin{
		"drafter": &rpc.Plugin{},
	}

	config := &plugin.ClientConfig{
		HandshakeConfig: rpc.HandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(cfg.cmdName()),
		Logger:          cfg.logger(),
	}

	return &engine{
		client: plugin.NewClient(config),
	}
}

func (e *engine) Dispense() (drafter.Drafter, error) {
	rpcClient, err := e.client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcClient.Dispense("drafter")
	if err != nil {
		return nil, err
	}

	return raw.(drafter.Drafter), nil
}

func (e *engine) Close() {
	e.client.Kill()
}
