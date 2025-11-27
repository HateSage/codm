package piscine

import "strconv"

func FromTo(from int, to int) string {
	// Check if any number is < 0 or > 99
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""
	start := from
	end := to
	step := 1

	// If from > to, we go backwards
	if from > to {
		step = -1
	}

	// Special case: from == to
	if from == to {
		if from < 10 {
			return "0" + strconv.Itoa(from) + "\n"
		}
		return strconv.Itoa(from) + "\n"
	}

	// Generate the range
	for i := start; ; i += step {
		// Format number: add leading zero if < 10
		if i < 10 {
			result += "0" + strconv.Itoa(i)
		} else {
			result += strconv.Itoa(i)
		}

		// Stop when we reach 'to'
		if i == end {
			break
		}

		// Add ", " separator
		result += ", "
	}

	// Add final newline
	result += "\n"
	return result
}