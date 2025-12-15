package functions

func FixArticles(str string) string {
	runes := []rune(str)
	result := ""
	i := 0
	for i < len(runes) {
		if runes[i] == 'a' || runes[i] == 'A' {
			if (i == 0 || !isLetterOrDigit(runes[i-1])) && (i == len(runes)-1 || !isLetterOrDigit(runes[i+1])) {
				j := i + 1
				for j < len(runes) && (runes[j] == ' ' || IsPunctuation(runes[j])) {
					j++
				}
				if IsVowelLetter(runes[j]) {
					if runes[i] == 'a' {
						result += string(runes[i]) + "n"
						i++
						continue
					} else if runes[i] == 'A' {
						result += string(runes[i]) + "N"
						i++
						continue
					}
				}
			}
			result += string(runes[i])
		} else {
			result += string(runes[i])
		}
		i++
	}
	return result
}
