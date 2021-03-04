package rpc

import (
	"bytes"
	"io"
	"net/rpc"

	"github.com/bukalapak/drafter-go"
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
	return c.call("Plugin.Parse", r, n)
}

func (c *Client) Check(r io.Reader, n drafter.Options) ([]byte, error) {
	return c.call("Plugin.Check", r, n)
}

func (c *Client) call(method string, r io.Reader, n drafter.Options) ([]byte, error) {
	var resp []byte

	b, err := readBytes(r)
	if err != nil {
		return nil, err
	}

	args := map[string]interface{}{
		"input":   b,
		"options": n,
	}

	if err = c.client.Call(method, args, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func readBytes(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)

	if _, err := io.Copy(buf, r); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
