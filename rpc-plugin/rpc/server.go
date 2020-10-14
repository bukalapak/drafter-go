package rpc

import (
	"bytes"

	"github.com/subosito/drafter-go"
)

type Server struct {
	Impl drafter.Drafter
}

func (v *Server) Version(args interface{}, resp *string) error {
	*resp = v.Impl.Version()
	return nil
}

func (v *Server) Parse(args map[string]interface{}, resp *string) error {
	b, err := v.Impl.Parse(bytes.NewReader(args["input"].([]byte)), args["options"].(drafter.Options))
	if err != nil {
		return err
	}

	*resp = string(b)
	return nil
}

func (v *Server) Check(args map[string]interface{}, resp *string) error {
	b, err := v.Impl.Check(bytes.NewReader(args["input"].([]byte)), args["options"].(drafter.Options))
	if err != nil {
		return err
	}

	*resp = string(b)
	return nil
}
