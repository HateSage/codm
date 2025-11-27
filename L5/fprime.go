package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		fmt.Println()
		return
	}

	factors := make([]int, 0)
	original := n

	// Check for factor 2
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Check for odd factors from 3 to sqrt(n)
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// If n is a prime number greater than 2
	if n > 1 {
		factors = append(factors, n)
	}

	// Print factors separated by *
	if len(factors) == 0 {
		fmt.Println(original)
	} else {
		for i, f := range factors {
			if i > 0 {
				fmt.Print("*")
			}
			fmt.Print(f)
		}
		fmt.Println()
	}
}
