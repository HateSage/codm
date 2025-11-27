package piscine

func Slice(a []string, nbrs ...int) []string {
	n := len(a)

	// Default start and end
	start := 0
	end := n

	// Handle variadic arguments
	if len(nbrs) >= 1 {
		start = nbrs[0]
	}
	if len(nbrs) >= 2 {
		end = nbrs[1]
	}

	// Convert negative indices
	if start < 0 {
		start = n + start
	}
	if end < 0 {
		end = n + end
	}

	// Clamp to valid bounds
	if start < 0 {
		start = 0
	}
	if end > n {
		end = n
	}

	// If start > end, return nil
	if start > end {
		return nil
	}

	// Return the slice
	return a[start:end]
}
