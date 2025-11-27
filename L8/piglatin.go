package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	word := os.Args[1]
	hasVowel := false
	for _, r := range word {
		if isVowel(r) {
			hasVowel = true
			break
		}
	}
	if !hasVowel {
		printStr("No vowels")
		z01.PrintRune('\n')
		return
	}
	firstVowel := -1
	for i, r := range word {
		if isVowel(r) {
			firstVowel = i
			break
		}
	}
	if firstVowel == 0 {
		printStr(word)
		printStr("ay")
	} else {
		printStr(word[firstVowel:])
		printStr(word[:firstVowel])
		printStr("ay")
	}
	z01.PrintRune('\n')
}

func isVowel(r rune) bool {
	r = toLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
