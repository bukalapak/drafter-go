# drafter-go

API blueprint parser for Go. It's a Go bindings for the [Drafter](https://github.com/apiaryio/drafter) library.


## Usage

This library can be consumed using two ways:

1. RPC-based plugin (recommended)
2. Go plugin

### RPC-based plugin

RPC-based plugin is based on Hashicorp's [go-plugin](https://github.com/hashicorp/go-plugin).

1. Download file drafter-rpc-* from releases page
2. Rename it as `drafter-rpc`
3. Add executable permission (chmod +x drafter.so)
4. Place it in the registered PATH directory
5. Run this file: `go run ./main.go`

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/subosito/drafter-go"
	rc "github.com/subosito/drafter-go/rpc-plugin/client"
)

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
```

### Go plugin

To consume this library as [go plugin](https://golang.org/pkg/plugin/). You need to:

1. Download file *.so from releases page
2. Rename it as `drafter.so`
3. Place it in the same directory with this file
4. Note: due to the go plugin limitation, we need to mimics GitHub build environment:
    ```sh
    export WORKDIR=/home/runner/work/drafter-go
    mkdir -p $WORKDIR
    git clone https://github.com/subosito/drafter-go.git $WORKDIR/drafter-go
    echo "replace github.com/subosito/drafter-go => $WORKDIR/drafter-go" >> go.mod
    ```
5. Run this file: `go run ./main.go`

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/subosito/drafter-go"
	rc "github.com/subosito/drafter-go/plugin/client"
)

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
```
