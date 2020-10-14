package drafter

import (
	"errors"
	"io"
)

type Drafter interface {
	Parse(r io.Reader, n Options) ([]byte, error)
	Check(r io.Reader, n Options) ([]byte, error)
	Version() string
}

const (
	YAML Format = iota
	JSON
)

var (
	ErrUnknown       = errors.New("unknown error")
	ErrInvalidInput  = errors.New("invalid input error")
	ErrInvalidOutput = errors.New("invalid output error")
)

type Format int

type Options struct {
	Format         Format
	SourceMaps     bool
	NameRequired   bool
	SkipBody       bool
	SkipBodySchema bool
}
