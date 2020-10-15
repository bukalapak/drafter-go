package rpc

import (
	"encoding/gob"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/subosito/drafter-go"
)

func init() {
	gob.Register(drafter.Options{})

}

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "DRAFTER_PLUGIN",
	MagicCookieValue: "drafter",
}

type Plugin struct {
	Impl drafter.Drafter
}

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &Server{Impl: p.Impl}, nil
}

func (p *Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &Client{client: c}, nil
}
