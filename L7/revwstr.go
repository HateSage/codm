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

	// Split into words (split by space)
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

	// Print words in reverse order
	for i := len(words) - 1; i >= 0; i-- {
		for _, r := range words[i] {
			z01.PrintRune(r)
		}
		if i > 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
