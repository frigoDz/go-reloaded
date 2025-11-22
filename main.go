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
func tokenize(str string) []string {
		var slice []string
		i := 0
		positon := 1
		var open int
		clos := 0
		trimed := ""
		for i < len(str) -1 {
				if str[i] == '\'' {
						position++
				}
				if str[i] == '\'' && position % 2 != 0 {
						open = i +1
						slice = append(slice,str[clos:open])
				}
				if str[i] == '\'' && position % 2 == 0 {
						clos = i
						trimed = strings.TrimSpace(str[open:clos])
						slice = append(slice,trimed)
				}
				if i == len(str) -1 {
						slice = append(slice ,str[clos:len(str)])
				}
		}
}

func applyNumberConversions(words []string) {
}


func applyCaseModifiers(words []string) []string {
}

func fixQuotes(str string) string {
}

func fixPunctuation(str string) string {
}

func fixArticles(words []string) []string{
}

func joinWords(words []string) string {
}
