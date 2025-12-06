package functions

import (
	"strconv"
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
func isApostropheRune(r rune) bool {
	if r == '\'' || r == '’' {
		return true
	}
	return false
}
func isWordApostrophe(runes []rune, i int) bool {
	if !(i > 0 && i < len(runes)-1 && unicode.IsLetter(runes[i+1]) && unicode.IsLetter(runes[i-1])) {
		return false
	}
	if !isApostropheRune(runes[i]) {
		return false
	}
	return true
}

func IsVowelLetter(c rune) bool {
	if c == 'a' || c == 'h' || c == 'o' || c == 'u' || c == 'i' || c == 'e' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'H' {
		return true
	}
	return false
}

func convertHex(str string) string {
	result := 0
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
			result = result*16 + value
		}
	} else {
		return str
	}
	return strconv.Itoa(result)
}

func convertBin(str string) string {
	result := 0
	if IsBin(str) {
		for i := 0; i < len(str); i++ {
			ch := str[i]
			var value int
			if ch == '0' {
				value = 0
			} else {
				value = 1
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
