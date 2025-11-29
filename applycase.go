package main

import (
		"strings"
		"unicode"
		"strconv"
)

func Applycase(str string) string {
	var strc string
	var res string
	start := -1
	var endd int
	var lastword string
	words := strings.Fields(str)
	for i, ch := range str {
		if ch == '(' {
			start = i
		}
		if start != -1 && ch == ')' {
			res += strings.TrimSpace(str[:start])
			strc = strings.TrimSpace(str[start+1:i])
			var num int
			if strings.Contains(strc, ",") {
				parts := strings.Split(strc, ",")
				if len(parts) == 2 {
					strc = strings.TrimSpace(parts[0])
					num, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
				}
			}
			lastword = foundword(res)
			switch strc {
			case "low":
				if num == 0 {
					lastword = strings.ToLower(lastword)
				} else {
					for j := len(words) - num; j < len(words); j++ {
						words[j] = strings.ToLower(words[j])
					}
				}
			case "up":
				if num == 0 {
					lastword = strings.ToUpper(lastword)
				} else {
					for j := len(words) - num; j < len(words); j++ {
						words[j] = strings.ToUpper(words[j])
					}
				}
			case "cap":
				if num == 0 {
					lastword = Capitalize(lastword)
				} else {
					for j := len(words) - num; j < len(words); j++ {
						words[j] = Capitalize(words[j])
					}
				}
			default:
				lastword += strc
			}
			res += lastword
			start = -1
		}
	}
	res += strings.Join(words, " ")[endd+1:]
	return res
}

func foundword(str string) string {
    res := ""
    for i := len(str)-1; i >= 0; i-- {
        if str[i] == ' ' {
            break
        } 
        res += string(str[i])
    }
    res = reversstr(res)
    return res
}
func reversstr(str string) string {
    res := ""
    for i := len(str)-1 ; i >= 0; i-- {
        res += string(str[i])
    }
    return res
}
func Capitalize (str string) string {
		str = strings.ToLower(str)
		res := ""
		i := 0 
		for i < len(str) {
				if unicode.IsLetter(rune(str[i])) {
						res += strings.ToUpper(string(str[i]))
						i++ 
						break
				}
				res += string(str[i])
				i++
		}
		res += str[i:]
		return res
}
