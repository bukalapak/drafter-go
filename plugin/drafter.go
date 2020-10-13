package main

import (
	"io"

	drafter "github.com/subosito/drafter-go"
	"github.com/subosito/drafter-go/adapter"
)

var Drafter engine

type engine struct{}

func (g engine) Parse(r io.Reader, n drafter.Options) ([]byte, error) {
	return adapter.Parse(r, n)
}

func (g engine) Check(r io.Reader, n drafter.Options) ([]byte, error) {
	return adapter.Check(r, n)
}

func (g engine) Version() string {
	return adapter.Version()
}
