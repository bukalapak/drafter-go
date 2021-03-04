package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bukalapak/drafter-go"
	rc "github.com/bukalapak/drafter-go/rpc-plugin/client"
)

/**
Usage:

1. Download file drafter-rpc-* from releases page
2. Rename it as `drafter-rpc`
3. Add executable permission (chmod +x drafter.so)
4. Place it in the registered PATH directory
5. Run this file: `go run ./main.go`
*/
func main() {
	client := rc.New(rc.Config{})
	defer client.Close()

	engine, err := client.Dispense()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	input := strings.NewReader("# API")
	options := drafter.Options{
		Format: drafter.JSON,
	}

	b, _ := engine.Parse(input, options)
	fmt.Printf("OUTPUT: %s\n", b)
}
