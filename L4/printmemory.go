package piscine

import "fmt"

func PrintMemory(arr [10]byte) {
	// Print hex values: 4 bytes per line, space-separated
	for i := 0; i < 10; i++ {
		// Print each byte as 2-digit hex
		fmt.Printf("%02x", arr[i])

		// Add space after each byte, except after 4th and 8th
		if i < 9 {
			fmt.Print(" ")
		}

		// Newline after 4th and 8th byte
		if i == 3 || i == 7 {
			fmt.Println()
		}
	}
	// Final newline after last group
	fmt.Println()

	// Print printable characters, dots for non-printable
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println() // final newline
}
