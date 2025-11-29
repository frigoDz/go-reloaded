package main

import "unicode"

func fixQuotes(str string) string {
    var slice string
    i := 0
    position := 0
    var open int
    clos := 0
    var valid bool
    for i < len(str) {
        if str[i] == '\'' {
            if !(i > 0 && i < len(str)-1 &&
                unicode.IsLetter(rune(str[i-1])) && unicode.IsLetter(rune(str[i+1]))) {
                position++
                valid = true
            } else {
                valid = false
            }
        }
        if str[i] == '\'' && position%2 != 0 && valid == true {
            before := str[clos:i]
			if len(before) > 0 && before[len(before)-1] != ' ' {
    			slice += before + " "
			} else {
    			slice += before
			}
			slice += "'"
            open = i + 1
            valid = false
        } else if str[i] == '\'' && position%2 == 0 && valid == true {
            clos = i
			trimed,newline := Trim(str[open:clos])
            slice += trimed
            slice += "'"
			if newline {
					slice += "\n"
			}
			clos = i + 1
			if clos < len(str) && str[clos] != ' ' {
    			slice += " "
			}
            valid = false
        }
        i++
    }
    slice += str[clos:]
    return slice
}

func Trim(str string) (string,bool) {
		strv2 := ""
		result := ""
		start := false
		start2 := false
		bol := false
		for _,ch := range str{
				if  ch != '\v' && ch != '\n' && ch != ' ' && ch != '\t' && ch != '\r' && ch != '\f' {
						start = true
				}
				if !start && (ch == '\v' || ch == '\n' || ch == ' ' || ch == '\t' || ch == '\r' || ch == '\f') {
						continue
				}
				strv2 += string(ch)
		}
		for i := len(strv2) -1; i >= 0; i-- {
				ch := strv2[i]
				if  ch != '\v' && ch != '\n' && ch != ' ' && ch != '\t' && ch != '\r' && ch != '\f' {
						start2 = true
				}
				if !start2 && (ch == '\v' || ch == '\n' || ch == ' ' || ch == '\t' || ch == '\r' || ch == '\f') {
						if ch == '\n' {
								bol = true
						}
						continue
				}
				result += string(ch)
		}
		return result,bol
}
