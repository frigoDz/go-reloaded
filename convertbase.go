package main

import (
	"strconv"
	"strings"
)

func applyNumberConversions(str string) string {
	words := strings.Fields(str)
	i := 0
	result := []string{}

	for i < len(words) {
		word := words[i]
		wr := []rune(word)
		punctuation := string(wr[len(wr):])
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			flag := string(wr[1 : len(wr)-1])
			if len(result) > 0 && (flag == "hex" || flag == "bin") && word == string(wr[:len(wr)]) {
				lastIdx := len(result) - 1
				if flag == "hex" {
					result[lastIdx] = convertHex(result[lastIdx])
				} else if flag == "bin" {
					result[lastIdx] = convertBin(result[lastIdx])
				}
				result[lastIdx] += punctuation
			} else {
				result = append(result, word)
			}
		} else {
			result = append(result, word)
		}
		i++
	}
	return strings.Join(result, " ")
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
