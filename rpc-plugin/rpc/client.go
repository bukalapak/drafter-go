package rpc

import (
	"io"
	"net/rpc"

	"github.com/subosito/drafter-go"
)

type Client struct {
	client *rpc.Client
}

func (c *Client) Version() string {
	var resp string

	if err := c.client.Call("Plugin.Version", new(interface{}), &resp); err != nil {
		return ""
	}

	return resp
}

func (c *Client) Parse(r io.Reader, n drafter.Options) ([]byte, error) {
	var resp string

	b, err := readBytes(r)
	if err != nil {
		return nil, err
	}

	args := map[string]interface{}{
		"input":   b,
		"options": n,
	}

	if err = c.client.Call("Plugin.Parse", args, &resp); err != nil {
		return nil, err
	}

	return []byte(resp), nil
}

func (c *Client) Check(r io.Reader, n drafter.Options) ([]byte, error) {
	var resp string

	b, err := readBytes(r)
	if err != nil {
		return nil, err
	}

	args := map[string]interface{}{
		"input":   b,
		"options": n,
	}

	if err = c.client.Call("Plugin.Check", args, &resp); err != nil {
		return nil, err
	}

	return []byte(resp), nil
}
