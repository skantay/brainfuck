# Brainfuck Interpreter

Brainfuck is an esoteric programming language created in 1993 by Urban Müller. It is known for its minimalistic and Turing-complete design, consisting of only eight simple commands, yet it can be used to create complex programs. This repository contains a Brainfuck interpreter implemented in Go.

## Getting Started

To run Brainfuck programs using this interpreter, follow these steps:

1. **Clone the Repository:**
   ```
   git clone https://github.com/username/brainfuck-interpreter.git
   cd brainfuck-interpreter
   ```

2. **Run Brainfuck Programs:**
   ```
   go run main.go "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
   ```

## Brainfuck Commands

- `>`: Increment the memory pointer.
- `<`: Decrement the memory pointer.
- `+`: Increment the byte at the memory pointer.
- `-`: Decrement the byte at the memory pointer.
- `.`: Output the byte at the memory pointer as a character.
- `[`: Jump forward to the corresponding `]` if the byte at the memory pointer is 0.
- `]`: Jump backward to the corresponding `[` if the byte at the memory pointer is nonzero.

## Example

Here's an example Brainfuck program that prints "Hello, World!":

```
go run . "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>." | cat -e
Hello World!$
```
```
go run . "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>." | cat -e
Hi$
```
```
go run . "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++." | cat -e
abc$
```
```
$ go run .
$
```
---