package main

import (
	"io"

	drafter "github.com/subosito/drafter-go"
)

var Drafter engine

type engine struct{}

func (g engine) Parse(r io.Reader, n drafter.Options) ([]byte, error) {
	return drafter.Parse(r, n)
}

func (g engine) Check(r io.Reader, n drafter.Options) ([]byte, error) {
	return drafter.Check(r, n)
}

func (g engine) Version() string {
	return drafter.Version()
}
