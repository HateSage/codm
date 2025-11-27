package piscine

func WordFlip(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	}
	words := []string{}
	word := ""
	for _, r := range str {
		if r == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	if len(words) == 0 {
		return "Invalid Output\n"
	}
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i > 0 {
			result += " "
		}
	}
	result += "\n"
	return result
}
