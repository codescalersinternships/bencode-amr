# Bencode

This project provides a Bencode decoder implemented in Go.

## Installation
To install the package, use `go get`:

```bash
go get github.com/codescalersinternships/bencode-amr

```

## Usage

### Import the bencode Parser package

```go
import bencode "github.com/codescalersinternships/bencode-amr/pkg"
```

### Decoding bencode Data
Use the Decode function
```go
value, remaining, err =  := bencode.Decode(<bencodeData>)
if err != nil {
    // Handle error
}
// Process the value

```
### Supported Data Types

- Integers
- Strings
- Lists
- Dictionaries

## Example 
Here's a simple example of how to use the RESP parser:
``` go
package main

import (
	"fmt"
	"log"

	bencode "github.com/codescalersinternships/bencode-amr/pkg"
)

func main() {
	encoded := []byte("d3:fooi42ee")
	value, remaining, err := bencode.Decode(encoded)
	if err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}
	fmt.Printf("Decoded Value: %+v\n", value)
	fmt.Printf("Remaining Bytes: %s\n", remaining)

```

## Testing
To run the tests for this package, use the following command:

```bash
make test
```
