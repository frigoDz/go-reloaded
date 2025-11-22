package main

import (
		"os"
		"strings"
		"fmt"
)
func main() {
		databyte,_ := os.ReadFile(os.Args[1])
		dataedit := string(databyte)
		if len(os.Args) > 3 {
				fmt.Print("Too many arguments")
				return
		}
		if len(os.Args) < 3 {
				fmt.Print("Not enough arguments")
				return
		}
		if len(datastr) == 0 {
				return
		}
		if os.Args[2] == "go.mod" || os.Args[2] == "go.sum" {
				return
		}
		dataedit = fixQuotes(dataedit)
		dataedit = fixPunctuation(dataedit)
		words := tokenize(dataedit)
		words = applyNumberConversions(words)
		words = applyCaseModifiers(words)
		words = FixArticle(words)
		dataedit = joinWords(words)
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
		for i < len(str)  {
				if str[i] == '\'' {
						if !(i > 0 && i < len(str)-1 &&
  ((str[i-1] >= 'a' && str[i-1] <= 'z') || (str[i-1] >= 'A' && str[i-1] <= 'Z')) &&
  ((str[i+1] >= 'a' && str[i+1] <= 'z') || (str[i+1] >= 'A' && str[i+1] <= 'Z'))) {
							position++
					}
				}
				if str[i] == '\'' && position % 2 != 0 {
						slice += str[clos:i]
						slice += "'"
						open = i+1
				}
				if str[i] == '\'' && position % 2 == 0{
						clos = i
						trimed = strings.TrimSpace(str[open:clos])
						slice += trimed
						slice += "'"
						clos = i + 1
				}
				i++
		}
		slice += str[clos:]
		return slice
}

func applyNumberConversions(words []string) {
}


func applyCaseModifiers(words []string) []string {
}

func tokenize(str string) string {
}

func fixPunctuation(str string) string {
}

func fixArticles(words []string) []string{
}

func joinWords(words []string) string {
}
