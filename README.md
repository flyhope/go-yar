# go-yar

Yar RPC Client for Go. This library provides a client implementation for the Yar RPC protocol in Go, supporting various serialization formats and robust logging.

## Features

*   **Yar RPC Protocol Support:** Implements the Yar RPC client protocol.
*   **Multiple Serialization Formats:** Supports JSON and MessagePack for data serialization.
*   **Flexible Logging:** Integrates with `logrus` for comprehensive logging capabilities.
*   **HTTP Transport:** Handles RPC requests over HTTP.

## Installation

To install `go-yar`, use `go get`:

```bash
go get github.com/flyhope/go-yar
```

## Usage

Here's a basic example of how to use the `go-yar` client:

```go
package main

import (
	"fmt"
	"github.com/flyhope/go-yar/client"
	"github.com/flyhope/go-yar/pack"
)

func main() {
	// Example: Create a new Yar client
	yarClient := client.NewClient("http://localhost:80/yar-server", client.WithPackager(pack.NewJsonPackager()))

	// Example: Call a remote method
	var result string
	err := yarClient.Call("Service.Method", []interface{}{"param1", 123}, &result)
	if err != nil {
		fmt.Printf("Error calling method: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
}
```

*Note: The example above is a placeholder. Please refer to the source code for detailed usage and available options.*

## Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.