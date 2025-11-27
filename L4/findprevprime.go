package piscine

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check divisors from 5 up to sqrt(n)
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// FindPrevPrime returns the largest prime <= nb
// Returns 0 if no such prime exists
func FindPrevPrime(nb int) int {
	if nb <= 2 {
		return 0
	}

	// Start from nb and go downwards
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}

	return 0
}