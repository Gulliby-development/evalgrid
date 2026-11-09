package main

// tokens normalizes text using ASCII-only rules before splitting on ledger blanks.
func tokens(text string) []string {
	cleaned := make([]byte, 0, len(text))
	for i := 0; i < len(text); i++ {
		b := text[i]
		switch {
		case b >= 'A' && b <= 'Z':
			b += 'a' - 'A'
		case b >= 33 && b <= 47,
			b >= 58 && b <= 64,
			b >= 91 && b <= 96,
			b >= 123 && b <= 126:
			continue
		}
		cleaned = append(cleaned, b)
	}

	var result []string
	start := -1
	for i := 0; i <= len(cleaned); i++ {
		if i == len(cleaned) || cleaned[i] == ' ' || cleaned[i] == '\t' {
			if start >= 0 {
				word := string(cleaned[start:i])
				if word != "a" && word != "an" && word != "the" {
					result = append(result, word)
				}
				start = -1
			}
		} else if start < 0 {
			start = i
		}
	}
	return result
}
