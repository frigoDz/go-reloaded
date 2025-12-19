package functions

import (
	"strconv"
	"strings"
	"unicode"
)

func IsPunctuation(c rune) bool {
	if c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';' {
		return true
	}
	return false
}

func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func IsVowelLetter(c rune) bool {
	if c == 'a' || c == 'h' || c == 'o' || c == 'u' || c == 'i' || c == 'e' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'H' {
		return true
	}
	return false
}

func convertHex(str string) string {
	result := 0
	maxint := 9223372036854775807
	if IsHex(str) {
		for i := 0; i < len(str); i++ {
			ch := str[i]
			var value int
			if ch >= '0' && ch <= '9' {
				value = int(ch - '0')
			} else if ch >= 'a' && ch <= 'f' {
				value = int(ch - 'a' + 10)
			} else if ch >= 'A' && ch <= 'F' {
				value = int(ch - 'A' + 10)
			}
			if result > (maxint-value)/16 {
				return str
			}
			result = result*16 + value
		}
	} else {
		return str
	}
	return strconv.Itoa(result)
}

func convertBin(str string) string {
	result := 0
	maxint := 9223372036854775807
	if IsBin(str) {
		for i := 0; i < len(str); i++ {
			ch := str[i]
			var value int
			if ch == '0' {
				value = 0
			} else {
				value = 1
			}
			if result > (maxint-value)/2 {
				return str
			}
			result = result*2 + value
		}
	} else {
		return str
	}
	return strconv.Itoa(result)
}

func IsBin(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, ch := range s {
		if ch != '0' && ch != '1' {
			return false
		}
	}
	return true
}

func IsHex(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, ch := range s {
		if !((ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'F') || (ch >= 'a' && ch <= 'f')) {
			return false
		}
	}
	return true
}
func Firstbyte(str string, targ rune) int {
	runes := []rune(str)
	for i, ch := range runes {
		if ch == targ {
			return i
		}
	}
	return -1
}

func Capitalize(str string) string {
	runes := []rune(str)
	runes = []rune(strings.ToLower(string(runes)))
	res := ""
	i := 0
	for i < len(runes) {
		if unicode.IsLetter(runes[i]) {
			res += strings.ToUpper(string(runes[i]))
			i++
			break
		}
		res += string(runes[i])
		i++
	}
	res += string(runes[i:])
	return res
}

func Matchp(runes []rune, start int) int {
	count := 1
	for i := start + 1; i < len(runes); i++ {
		if runes[i] == '(' {
			count++
		} else if runes[i] == ')' {
			count--
			if count == 0 {
				return i
			}
		}
	}
	return -1
}

func splitWords(str string) []string {
	words := []string{}
	currentWord := ""
	for _, ch := range str {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(ch)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}
	return words
}

func rebuildWithWhitespace(original string, transformedWords []string) string {
	result := ""
	wordIdx := 0
	currentWord := ""

	for _, ch := range original {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if currentWord != "" {
				if wordIdx < len(transformedWords) {
					result += transformedWords[wordIdx]
					wordIdx++
				} else {
					result += currentWord
				}
				currentWord = ""
			}
			result += string(ch)
		} else {
			currentWord += string(ch)
		}
	}

	if currentWord != "" {
		if wordIdx < len(transformedWords) {
			result += transformedWords[wordIdx]
		} else {
			result += currentWord
		}
	}

	return result
}

func CheckPun(s string) bool {
	for _, ch := range s {
		if ch == ' ' || ch == '\t' {
			continue
		}
		return IsPunctuation(ch)
	}
	return false
}
