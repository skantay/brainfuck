package brainfuck

import (
	"bufio"
	"errors"
	"io"
)

func BrainfuckExtended1(file io.Reader) (string, error) {
	scanner := bufio.NewScanner(file)
	var code string

	for scanner.Scan() {
		code += scanner.Text()
	}
	return interpretBrainfuckExtended1([]rune(code))
}

// _
func interpretBrainfuckExtended1(code []rune) (string, error) {
	memory := make([]byte, 30000)
	var pointer int
	output := make([]byte, 0)
	var storage byte

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
			output = append(output, memory[pointer])
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
		case '@':
			return string(output), nil
		case '$':
			storage = memory[pointer]
		case '!':
			memory[pointer] = storage
		case '}':
			memory[pointer] >>= 1
		case '{':
			memory[pointer] <<= 1
		case '~':
			memory[pointer] = ^memory[pointer]
		case '^':
			memory[pointer] ^= storage
		case '&':
			memory[pointer] &= storage
		case '|':
			memory[pointer] |= storage
		default:
			return "", errors.New("compiler error: unknown char")
		}
	}

	return string(output), nil
}
