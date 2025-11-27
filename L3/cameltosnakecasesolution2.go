package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	x := func(c byte) bool {
		return c >= 'a' && c <= 'z'
	}
	y := func(c byte) bool {
		return c >= 'A' && c <= 'Z'
	}

	if s[0] < 'A' || s[0] > 'Z' && s[0] < 'a' || s[0] > 'z' {
		return s
	}
	for i := 0; i < len(s); i++ {
		if !y(s[i]) && !x(s[i]) {
			return s
		}
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 && y(s[i]) {
			return s
		}
		if y(s[i]) && y(s[i+1]) {
			return s
		}
	}

	var a []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		if y(c) {
			if i > 0 {
				a = append(a, '_')
			}
		}
		a = append(a, c)
	}
	return string(a)
}