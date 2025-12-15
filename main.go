package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: Invalid number of argument")
		return
	}
	if !(strings.HasSuffix(os.Args[2], ".txt")) {
		fmt.Println("Error: Incorrect file type")
		return
	}
	info, er := os.Stat(os.Args[1])
	if er != nil {
		fmt.Println("Error:", er)
		return
	}
	if info.Size() > 1_000_000 {
		fmt.Println("Error: file is too large")
		return
	}
	databyte, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	dataedit := string(databyte)
	if len(dataedit) == 0 {
		return
	}
	dataedit = functions.Applycase(dataedit)
	dataedit = functions.FixArticles(dataedit)
	dataedit = functions.FixPunctuation(dataedit)
	dataedit = functions.FixQuotes(dataedit)
	res, errr := os.Create(os.Args[2])
	if errr != nil {
		fmt.Println("Error", errr)
		return
	}
	res.WriteString(dataedit)
	res.Close()
}
