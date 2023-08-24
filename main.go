package main

import "fmt"

func Balanced(input string) bool {
	bracketStack := make([]rune, 0)

	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range input {
		switch char {
		case '(', '[', '{':
			bracketStack = append(bracketStack, char)
		case ')', ']', '}':
			if len(bracketStack) == 0 || bracketStack[len(bracketStack)-1] != bracketPairs[char] {
				return false
			}
			bracketStack = bracketStack[:len(bracketStack)-1]
		}
	}

	return len(bracketStack) == 0
}

func main() {
	input1 := "{()}"
	if Balanced(input1) {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}

	input2 := "{()}(])"
	if Balanced(input2) {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}
}
