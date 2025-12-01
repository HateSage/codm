package main

import (
	"fmt"
	"os"
)

// minimal Atoi for non-negative integers
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	if os.Args[1] != "-c" {
		os.Exit(1)
	}

	n, ok := atoi(os.Args[2])
	if !ok || n < 0 {
		os.Exit(1)
	}

	files := os.Args[3:]
	errorOccurred := false
	hasOutput := false

	for _, filename := range files {

		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			errorOccurred = true
			hasOutput = true
			continue
		}

		// Print blank line BEFORE header if ANY prior output (error or content)
		if hasOutput {
			fmt.Println()
		}

		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", filename)
		}

		hasOutput = true

		if n > len(data) {
			fmt.Print(string(data))
		} else {
			fmt.Print(string(data[len(data)-n:]))
		}
	}

	if errorOccurred {
		os.Exit(1)
	}
}
