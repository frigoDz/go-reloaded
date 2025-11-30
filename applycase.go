package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Applycase(str string) string {
	result := ""
	i := 0

	for i < len(str) {
		start := Firstbyte(str[i:], '(')
		if start == -1 {
			result += str[i:]
			break
		}
		start += i
		end := Firstbyte(str[start:], ')')
		if end == -1 {
			result += str[i : start+1]
			i = start + 1
			continue
		}
		end += start
		command := strings.TrimSpace(str[start+1 : end])
		var cmd string
		var num int
		validCommand := false

		if strings.Contains(command, ",") {
			parts := strings.Split(command, ",")
			if len(parts) == 2 {
				cmd = strings.TrimSpace(parts[0])
				numStr := strings.TrimSpace(parts[1])
				var err error
				num, err = strconv.Atoi(numStr)
				if err == nil && (cmd == "low" || cmd == "up" || cmd == "cap") {
					validCommand = true
				}
			}
		} else {
			cmd = command
			num = 1
			if cmd == "low" || cmd == "up" || cmd == "cap" {
				validCommand = true
			}
		}
		if !validCommand {
			result += str[i : start+1]
			i = start + 1
			continue
		}
		textBefore := str[i:start]
		allText := result + textBefore
		words := strings.Fields(allText)
		if len(words) >= num {
			for j := len(words) - num; j < len(words); j++ {
				switch cmd {
				case "low":
					words[j] = strings.ToLower(words[j])
				case "up":
					words[j] = strings.ToUpper(words[j])
				case "cap":
					words[j] = Capitalize(words[j])
				}
			}
			result = strings.Join(words, " ")
		} else {
			result += textBefore
		}
		i = end + 1
		if i < len(str) && str[i] == ' ' {
			i++
			if i < len(str) {
				result += " "
			}
		}
	}
	return result
}

func Firstbyte(str string, targ rune) int {
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == targ {
			return i
		}
	}
	return -1
}

func Capitalize(str string) string {
	str = strings.ToLower(str)
	res := ""
	i := 0
	for i < len(str) {
		if unicode.IsLetter(rune(str[i])) {
			res += strings.ToUpper(string(str[i]))
			i++
			break
		}
		res += string(str[i])
		i++
	}
	res += str[i:]
	return res
}
