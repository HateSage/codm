package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	len1 := len(slice1)
	len2 := len(slice2)

	// Determine which slice is longer (or equal)
	if len1 > len2 {
		// slice1 longer → start with slice1 reversed
		for i := len1 - 1; i >= 0; i-- {
			result = append(result, slice1[i])
			if i < len2 {
				result = append(result, slice2[len2-1-i])
			}
		}
	} else if len2 > len1 {
		// slice2 longer → start with slice2 reversed
		for i := len2 - 1; i >= 0; i-- {
			result = append(result, slice2[i])
			if i < len1 {
				result = append(result, slice1[len1-1-i])
			}
		}
	} else {
		// equal length → start with slice1 reversed
		for i := len1 - 1; i >= 0; i-- {
			result = append(result, slice1[i])
			result = append(result, slice2[i])
		}
	}

	return result
}
