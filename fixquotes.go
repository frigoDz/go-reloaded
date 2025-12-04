package main

import "unicode"

func fixQuotes(str string) string {
	runes := []rune(str)
	var slice string
	i := 0
	position := 0
	var open int
	clos := 0
	var valid bool
	for i < len(runes) {
		if runes[i] == '\'' {
			if !(i > 0 && i < len(runes)-1 &&
				unicode.IsLetter(rune(runes[i-1])) && unicode.IsLetter(rune(runes[i+1]))) {
				position++
				valid = true
			} else {
				valid = false
			}
		}
		if runes[i] == '\'' && position%2 != 0 && valid == true {
			before := runes[clos:i]
			if len(before) > 0 && before[len(before)-1] != ' ' {
				slice += string(before) + " "
			} else {
				slice += string(before)
			}
			open = i + 1
			clos = i
			valid = false
		} else if runes[i] == '\'' && position%2 == 0 && valid == true {
			clos = i
			trimed, newline := Trim(string(runes[open:clos]))
			slice += "'" + trimed + "'"
			if newline {
				slice += "\n"
			}
			clos = i + 1
			if clos < len(runes) && runes[clos] != ' ' {
				slice += " "
			}
			valid = false
		}
		i++
	}
	slice += string(runes[clos:])
	return slice
}

func Trim(str string) (string, bool) {
	runes := []rune(str)
	strv2 := ""
	start := false
	bol := false
	for _, ch := range runes {
		if ch != '\v' && ch != '\n' && ch != ' ' && ch != '\t' && ch != '\r' && ch != '\f' {
			start = true
		}
		if !start && (ch == '\v' || ch == '\n' || ch == ' ' || ch == '\t' || ch == '\r' || ch == '\f') {
			continue
		}
		strv2 += string(ch)
	}
	result := ""
	start2 := false
	for i := len(strv2) - 1; i >= 0; i-- {
		ch := rune(strv2[i])
		if ch != '\v' && ch != '\n' && ch != ' ' && ch != '\t' && ch != '\r' && ch != '\f' {
			start2 = true
		}
		if !start2 && (ch == '\v' || ch == '\n' || ch == ' ' || ch == '\t' || ch == '\r' || ch == '\f') {
			if ch == '\n' {
				bol = true
			}
			continue
		}
		result = string(ch) + result
	}
	return result, bol
}
