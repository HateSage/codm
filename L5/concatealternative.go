package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	len1 := len(slice1)
	len2 := len(slice2)

	// Determine which slice is longer (or equal)
	if len1 >= len2 {
		// slice1 is longer or equal → start with slice1
		for i := 0; i < len1; i++ {
			result = append(result, slice1[i])
			if i < len2 {
				result = append(result, slice2[i])
			}
		}
	} else {
		// slice2 is longer → start with slice2
		for i := 0; i < len2; i++ {
			result = append(result, slice2[i])
			if i < len1 {
				result = append(result, slice1[i])
			}
		}
	}

	return result
}
