package main

import (
	"fmt"
	"os"
)

func main() {
	databyte, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
		return
	}
	dataedit := string(databyte)
	if len(os.Args) < 3 {
		fmt.Print("Not enough arguments")
		return
	}
	if len(os.Args) > 3 {
		fmt.Print("Too many arguments")
		return
	}
	if len(dataedit) == 0 {
		return
	}
	dataedit = fixArticles(dataedit)
	dataedit = fixPunctuation(dataedit)
	dataedit = fixQuotes(dataedit)
	dataedit = applyNumberConversions(dataedit)
	dataedit = Applycase(dataedit)
	res, _ := os.Create(os.Args[2])
	res.WriteString(dataedit)
	res.Close()
}
