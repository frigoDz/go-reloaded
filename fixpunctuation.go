package main

import "strings"

func fixPunctuation(str string) string {
		slice := ""
		for i := 0; i < len(str); i++ {
				if IsPunctuation(rune(str[i])) {
						slice = strings.TrimRight(slice, " ")
						j := i
						for j + 1 < len(str) && IsPunctuation(rune(str[j+1])) {
								j++
						}
						slice += string(str[i:j+1])
						if j < len(str)-1 && str[j+1] != ' ' {
								slice += " "
						}
						i = j
				} else {
						slice += string(str[i])
				}
		}
		return slice
}
