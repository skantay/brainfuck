package brainfuck

import (
	"fmt"
	"os"
)

func Brainfuck(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	var res string

	for scanner.Scan() {
		res += scanner.Text()
	}

	return interpretBrainfuck(res)
}

func interpretBrainfuck(code string) string {
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

	return output
}

