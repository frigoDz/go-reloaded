package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Applycase(str string) string {
	runes := []rune(str)
	result := ""
	i := 0
	for i < len(runes) {
		start := Firstbyte(string(runes[i:]), '(')
		if start == -1 {
			result += string(runes[i:])
			break
		}
		start += i
		if start > 0 && runes[start-1] != ' ' && runes[start-1] != '\t' && runes[start-1] != '\n' {
				result += string(runes[i:start+1])
				i = start + 1
				continue
		}
		end := Firstbyte(string(runes[start:]), ')')
		if end == -1 {
			result += string(runes[i : start+1])
			i = start + 1
			continue
		}
		end += start
		if end < len(runes)-1 && runes[end+1] != ' ' && runes[end+1] != '\t' && runes[end+1] != '\n' {
				result += string(runes[i:start+1])
				i = start + 1
				continue
		}
		command := strings.TrimSpace(string(runes[start+1 : end]))
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
			result += string(runes[i : start+1])
			i = start + 1
			continue
		}
		textBefore := string(runes[i:start])
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
		if i < len(runes) && runes[i] == ' ' {
			i++
			if i < len(runes) {
				result += " "
			}
		}
	}
	return result
}

func Firstbyte(str string, targ rune) int {
	runes := []rune(str)
	for i,ch := range runes {
		if ch == targ {
			return i
		}
	}
	return -1
}

func Capitalize(str string) string {
	runes := []rune(str)
	runes = []rune(strings.ToLower(string(runes)))
	res := ""
	i := 0
	for i < len(runes) {
		if unicode.IsLetter(runes[i]) {
			res += strings.ToUpper(string(runes[i]))
			i++
			break
		}
		res += string(runes[i])
		i++
	}
	res += string(runes[i:])
	return res
}
