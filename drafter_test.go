package drafter_test

import (
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	drafter "github.com/subosito/drafter-go"
)

func TestDrafter_Parse(t *testing.T) {
	c := qt.New(t)
	s := strings.NewReader("# API")

	b, err := drafter.Parse(s, drafter.Options{
		Format: drafter.JSON,
	})

	c.Assert(err, qt.IsNil)
	c.Assert(string(b), qt.Contains, "API")
}

func TestDrafter_Check(t *testing.T) {
	c := qt.New(t)
	n := drafter.Options{}

	s := strings.NewReader("# API")
	b, err := drafter.Check(s, n)
	c.Assert(err, qt.IsNil)
	c.Assert(string(b), qt.Equals, "")

	s = strings.NewReader("# API\n## Data Structures\n### Hello-World (object)\n+ foo: bar (string, required)")
	b, err = drafter.Check(s, n)
	c.Assert(err, qt.IsNil)
	c.Assert(string(b), qt.Contains, "please escape the name of the data structure using backticks")
}

func TestDrafter_Version(t *testing.T) {
	c := qt.New(t)
	c.Assert(drafter.Version(), qt.Equals, "v5.0.0")
}
