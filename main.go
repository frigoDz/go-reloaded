package main

import (
	"fmt"
	"os"
)

func main() {
	databyte, _ := os.ReadFile(os.Args[1])
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
	dataedit = fixArticles(dataedit)
	dataedit = Applycase(dataedit)
	res, _ := os.Create(os.Args[2])
	res.WriteString(dataedit)
	res.Close()
}
