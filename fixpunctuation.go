package main

import "strings"

func fixPunctuation(str string) string {
		runes := []rune(str)
		slice := ""
		for i := 0; i < len(runes); i++ {
			if IsPunctuation(runes[i]) {
				slice = strings.TrimRight(slice, " ")
				j := i
				for j+1 < len(runes) && IsPunctuation(runes[j+1]) {
					j++
				}
				slice += string(runes[i:j+1])
				if j < len(runes)-1 && runes[j+1] != ' ' {
					slice += " "
				}
				i = j
			} else {
				slice += string(runes[i])
			}
		}
		return slice
}
