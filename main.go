package main

import (
		"os"
		"strings"
		"fmt"
		"unicode"
)
func main() {
		databyte,_ := os.ReadFile(os.Args[1])
		dataedit := string(databyte)
		if len(os.Args) < 3 {
				fmt.Print("Too many arguments")
				return
		}
		if len(os.Args) > 3 {
				fmt.Print("Not enough arguments")
				return
		}
		if len(dataedit) == 0 {
				return
		}
		if os.Args[2] == "go.mod" || os.Args[2] == "go.sum" {
				return
		}
		dataedit = fixQuotes(dataedit)
		dataedit = fixPunctuation(dataedit)
		dataedit = applyNumberConversions(dataedit)
//		dataedit = applyCaseModifiers(dataedit)
		dataedit = fixArticles(dataedit)
		res,_ := os.Create(os.Args[2])
		res.WriteString(dataedit)
		res.Close()
}
func fixQuotes(str string) string {
    var slice string
    i := 0
    position := 0
    var open int
    clos := 0
    trimed := ""
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
            trimed = strings.TrimSpace(str[open:clos])
            slice += trimed
            slice += "'"
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
func applyNumberConversions(sentence string) string{
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		w := words[i]
		if strings.HasPrefix(w, "(") && strings.HasSuffix(w, ")") {
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

/*
func applyCaseModifiers(str string) string {
}*/


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
func fixArticles(str string) string {
		runes := []rune(str)
		result := ""
		i := 0
		for i < len(runes) {
				if runes[i] == 'a' || runes[i] == 'A' {
						if (i == 0 || !isLetterOrDigit(runes[i-1])) && (i == len(runes)-1 || !isLetterOrDigit(runes[i+1])) {
								j := i + 1
								for j < len(runes) && (runes[j] == ' ' || IsPunctuation(runes[j])) {
										j++
								}
								if IsVowelLetter(runes[j]) {
										result += string(runes[i]) + "n"
										i++
										continue
								}
						}
						result += string(runes[i])
				} else {
						result += string(runes[i])
				}
				i++
		}
		return result
}
