# Brainfuck Interpreter

Brainfuck is an esoteric programming language created in 1993 by Urban Müller. It is known for its minimalistic and Turing-complete design, consisting of only eight simple commands, yet it can be used to create complex programs. This repository contains a Brainfuck interpreter implemented in Go.

## Getting Started

To run Brainfuck programs using this interpreter, follow these steps:

1. **Go get the package:**
   ```
   go get github.com/skantay/brainfuck@1.0.0
   ```

2. **Run Brainfuck Programs:**
   ```go
   package main

   import "github.com/skantay/brainfuck"

   func main() {
      file, _ := os.Open("file.brainfuck")
      result := brainfuck.Brainfuck(file)
   }
   ```

## Brainfuck language syntax

- '>' increment the pointer
- '<' decrement the pointer
- '+' increment the pointed byte
- '-' decrement the pointed byte
- '.' print the pointed byte on standard output
- '[' go to the matching ']' if the pointed byte is 0 (loop start)
- ']' go to the matching '[' if the pointed byte is not 0 (loop end)

## Example

Here's an example Brainfuck program that prints "Hello, World!":

```
++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
```

```
Hello World!$
```
---
