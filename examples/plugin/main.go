package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bukalapak/drafter-go"
	rc "github.com/bukalapak/drafter-go/plugin/client"
)

/**
Usage:

1. Download file *.so from releases page
2. Rename it as `drafter.so`
3. Place it in the same directory with this file
4. Note: due to the go plugin limitation, there are additional steps to mimics plugin build environment:
    1. export WORKDIR=/home/runner/work/drafter-go
    2. mkdir -p $WORKDIR
    3. git clone https://github.com/bukalapak/drafter-go.git $WORKDIR/drafter-go
    4. echo "replace github.com/bukalapak/drafter-go => $WORKDIR/drafter-go" >> go.mod
5. Run this file: `go run ./main.go`
*/
func main() {
	engine, err := rc.New(rc.Config{})
	if err != nil {
		log.Fatalf("Error: %s\n", err)
		os.Exit(1)
	}

	input := strings.NewReader("# API")
	options := drafter.Options{
		Format: drafter.JSON,
	}

	b, _ := engine.Parse(input, options)
	fmt.Printf("OUTPUT: %s\n", b)
}
