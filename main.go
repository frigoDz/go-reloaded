package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
	"strings"
)

func main() {
	if !(strings.HasSuffix(os.Args[2], ".txt")) {
		fmt.Print("Incorrect file type")
		return
	}
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
	dataedit = functions.Applycase(dataedit)
	dataedit = functions.FixArticles(dataedit)
	dataedit = functions.FixQuotes(dataedit)
	dataedit = functions.FixPunctuation(dataedit)
	res, errr := os.Create(os.Args[2])
	if errr != nil {
		fmt.Print(errr)
		return
	}
	res.WriteString(dataedit)
	res.Close()
}
