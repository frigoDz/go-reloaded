package main

import (
		"strconv"
		"strings"
)

func applyNumberConversions(sentence string) string{
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		w := words[i]
		if strings.HasPrefix(w, "(") && strings.HasSuffix(strings.TrimRight(w, ".,!?;:"), ")") {
			inner := strings.TrimSpace(w[1 : len(w)-1])
			if inner == "hex" && i > 0 {
				words[i-1] = convertHex(words[i-1])
				words[i] = ""
			}
			if inner == "bin" && i > 0 {
				words[i-1] = convertBin(words[i-1])
				words[i] = ""
			}
		}
	}
	result := []string{}
	for _, w := range words {
		if w != "" {
			result = append(result, w)
		}
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
						result = result * 16 + value
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
		}else{
				return str
		}
		return strconv.Itoa(result)
}

func IsBin(s string) bool {
		for _,ch := range s {
				if ch != '0' && ch != '1' {
						return false
				}
		}
		return true
}

func IsHex(s string) bool {
		for _,ch := range s {
				if !((ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'F') ||  (ch >= 'a' && ch <= 'f')) {
						return false
				}
		}
		return true
}

