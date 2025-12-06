package functions

import (
	"strings"
	"unicode"
)

func FixQuotes(str string) string {
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
			clos = i + 1
			valid = false
		} else if runes[i] == '\'' && position%2 == 0 && valid == true {
			clos = i
			newline := foundNl(string(runes[open:clos]))
			trimed := strings.TrimSpace(string(runes[open:clos]))
			slice += "'" + trimed + "'"
			if newline {
				slice += "\n"
			}
			clos = i + 1
			if clos < len(runes) && (runes[clos] == ' ' || runes[clos] == '\t') {
				clos++
			}
			if newline && clos < len(runes) && runes[clos] == '\n' {
				clos++
				for clos < len(runes) && (runes[clos] == ' ' || runes[clos] == '\t') {
					clos++
				}
			} else if !newline && clos < len(runes) && runes[clos] != ' ' && runes[clos] != '\n' {
				slice += " "
			}
			valid = false
		}
		i++
	}
	slice += string(runes[clos:])
	return slice
}

func foundNl(str string) bool {
	res := false
	for _, ch := range str {
		if ch == '\n' {
			res = true
		}
	}
	return res
}
