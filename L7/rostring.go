// rostring.go
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Must have exactly one argument
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	str := os.Args[1]

	// Extract words, ignoring extra spaces
	words := []string{}
	word := ""
	for _, r := range str {
		if r == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		words = append(words, word)
	}

	// If no words → just newline
	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Print all words except first
	for i := 1; i < len(words); i++ {
		for _, r := range words[i] {
			z01.PrintRune(r)
		}
		if i < len(words)-1 {
			z01.PrintRune(' ')
		}
	}

	// Print first word at the end
	if len(words) > 1 {
		z01.PrintRune(' ')
	}
	for _, r := range words[0] {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}
