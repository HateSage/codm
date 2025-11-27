package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := make(map[rune]bool) // to track printed characters (avoid doubles)

	// Loop through s1
	for _, char := range s1 {
		// Check if char exists in s2
		found := false
		for _, c2 := range s2 {
			if c2 == char {
				found = true
				break
			}
		}

		// If found in s2 and not printed yet → print it
		if found && !seen[char] {
			seen[char] = true
			z01.PrintRune(char)
		}
	}

	z01.PrintRune('\n')
}
