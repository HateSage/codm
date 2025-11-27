package piscine

func WeAreUnique(str1, str2 string) int {
	// Special case: both strings empty → return -1
	if str1 == "" && str2 == "" {
		return -1
	}

	// Map to track which characters appear in str1
	inStr1 := make(map[rune]bool)

	// Map to track which characters appear in str2
	inStr2 := make(map[rune]bool)

	// First pass: mark all characters that appear in str1
	for _, r := range str1 {
		inStr1[r] = true
	}

	// Second pass: mark all characters that appear in str2
	for _, r := range str2 {
		inStr2[r] = true
	}

	// Count characters that are in str1 but NOT in str2
	count := 0
	for r := range inStr1 {
		if !inStr2[r] { // r is in str1 but NOT in str2 → unique
			count++
		}
	}

	// Count characters that are in str2 but NOT in str1
	for r := range inStr2 {
		if !inStr1[r] { // r is in str2 but NOT in str1 → unique
			count++
		}
	}

	// Return total number of unique characters
	return count
}
