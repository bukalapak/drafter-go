package client

import (
	"plugin"

	"github.com/subosito/drafter-go"
)

type Config struct {
	Name string
}

func (c Config) pluginName() string {
	if c.Name == "" {
		return "drafter.so"
	}

	return c.Name
}

func New(cfg Config) (drafter.Drafter, error) {
	plug, err := plugin.Open(cfg.pluginName())
	if err != nil {
		return nil, err
	}

	symDrafter, err := plug.Lookup("Drafter")
	if err != nil {
		return nil, err
	}

	engine, ok := symDrafter.(drafter.Drafter)
	if !ok {
		return nil, err
	}

	return engine, nil
}
