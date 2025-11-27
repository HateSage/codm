package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	newWord := true // We're at the start of a word

	for _, char := range s {
		if newWord {
			// First character of a word must be uppercase OR non-lowercase letter
			if char >= 'a' && char <= 'z' {
				return false // lowercase letter at start of word → invalid
			}
			newWord = false // we've processed the first char of this word
		}

		// Space means next character starts a new word
		if char == ' ' {
			newWord = true
		}
	}

	return true
}