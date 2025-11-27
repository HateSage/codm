package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	if !isValidCamelCase(s) {
		return s
	}

	result := ""
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(char)
		} else {
			result += string(char)
		}
	}

	return result
}

func isValidCamelCase(s string) bool {
	for i, char := range s {
		if char >= '0' && char <= '9' {
			return false
		}

		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return false
		}

		if char >= 'A' && char <= 'Z' {
			if i == len(s)-1 {
				return false
			}

			if i > 0 {
				prevChar := rune(s[i-1])
				if prevChar >= 'A' && prevChar <= 'Z' {
					return false
				}
			}
		}
	}

	return true
}
