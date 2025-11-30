package main

import "unicode"

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
