package adapter

import (
	"io"

	drafter "github.com/bukalapak/drafter-go"
	"github.com/bukalapak/drafter-go/adapter"
)

type Adapter struct{}

func (g Adapter) Parse(r io.Reader, n drafter.Options) ([]byte, error) {
	return adapter.Parse(r, n)
}

func (g Adapter) Check(r io.Reader, n drafter.Options) ([]byte, error) {
	return adapter.Check(r, n)
}

func (g Adapter) Version() string {
	return adapter.Version()
}

func New() Adapter {
	return Adapter{}
}
