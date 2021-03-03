package rpc

import (
	"bytes"

	"github.com/bukalapak/drafter-go"
)

type Server struct {
	Impl drafter.Drafter
}

func (v *Server) Version(args interface{}, resp *string) error {
	*resp = v.Impl.Version()
	return nil
}

func (v *Server) Parse(args map[string]interface{}, resp *[]byte) error {
	input, options := v.parseArgs(args)
	b, err := v.Impl.Parse(input, options)
	if err != nil {
		return err
	}

	*resp = b
	return nil
}

func (v *Server) Check(args map[string]interface{}, resp *[]byte) error {
	input, options := v.parseArgs(args)
	b, err := v.Impl.Check(input, options)
	if err != nil {
		return err
	}

	*resp = b
	return nil
}

func (v *Server) parseArgs(args map[string]interface{}) (*bytes.Reader, drafter.Options) {
	return bytes.NewReader(args["input"].([]byte)), args["options"].(drafter.Options)
}
