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
   go run main.go your-brainfuck-program.bf
   ```
   Replace `your-brainfuck-program.bf` with the path to your Brainfuck source file.

## Brainfuck Commands

- `>`: Increment the memory pointer.
- `<`: Decrement the memory pointer.
- `+`: Increment the byte at the memory pointer.
- `-`: Decrement the byte at the memory pointer.
- `.`: Output the byte at the memory pointer as a character.
- `,`: Input a character and store it in the byte at the memory pointer.
- `[`: Jump forward to the corresponding `]` if the byte at the memory pointer is 0.
- `]`: Jump backward to the corresponding `[` if the byte at the memory pointer is nonzero.

## Example

Here's an example Brainfuck program that prints "Hello, World!":

```brainfuck
++++++++[>++++[>++>+++>+++>+<<<<-]>+.-.>++.+++.>++.<<-----.>-.>.+++.------.>+.>.]
```

To run this program using the interpreter:

```
go run main.go hello-world.bf
```

## Contributing

If you want to contribute to this project and make it better, your help is very welcome. Create a fork of the project, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---