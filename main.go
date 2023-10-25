package main

import (
	"fmt"
	"os"
)

func interpretBrainfuck(code string) {
	const memorySize = 30000
	var memory [memorySize]byte
	var pointer int
	var output string

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			output += string(memory[pointer])
		case '[':
			if memory[pointer] == 0 {
				loopCount := 1
				for loopCount > 0 {
					i++
					if code[i] == '[' {
						loopCount++
					} else if code[i] == ']' {
						loopCount--
					}
				}
			}
		case ']':
			if memory[pointer] != 0 {
				loopCount := 1
				for loopCount > 0 {
					i--
					if code[i] == ']' {
						loopCount++
					} else if code[i] == '[' {
						loopCount--
					}
				}
			}
		}
	}

	fmt.Print(output)
}

func main() {
	code := os.Args[1:]
	if len(code) == 1 {
		interpretBrainfuck(code[0])
	}
}
