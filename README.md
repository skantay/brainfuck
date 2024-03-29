# Go Eso Compiler


The Go Eso Compiler is a package that provides functionality to compile programs written in [Eso](https://esolangs.org/wiki/Main_Page) programming languages.

## Getting started

To use this package, you need to have Go installed.

Go get the package:
```bash
go get github.com/skantay/brainfuck
```

## Example

Here's a simple example demonstrating how to use the package:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/skantay/go-eso-compiler"
)

func main() {
	file, err := os.Open("program.bf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := brainfuck.Brainfuck(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Compiled output:", output)
}
```

## Supported Languages

- [Brainfuck](https://esolangs.org/wiki/Brainfuck)

- [Brainfuck Extended Type 1](https://esolangs.org/wiki/Extended_Brainfuck#Extended_Type_I)

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.
