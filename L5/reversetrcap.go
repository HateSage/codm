package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		length := 0
		for range arg {
			length++
		}

		wordEnd := false
		for j, char := range arg {
			if isAlpha(char) {
				if wordEnd || j == length-1 {
					// Last letter of word or last char → UPPERCASE
					z01.PrintRune(toUpper(char))
				} else {
					// Not last letter → lowercase
					z01.PrintRune(toLower(char))
				}
				wordEnd = false
			} else {
				// Non-alpha: space, punctuation, number → print as is
				z01.PrintRune(char)
				wordEnd = true
			}
		}
		z01.PrintRune('\n')
	}
}
