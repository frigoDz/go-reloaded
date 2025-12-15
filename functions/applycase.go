package functions

import (
	"strconv"
	"strings"
)

func Applycase(str string) string {
	prevresult := ""
	result := str
	for result != prevresult {
		prevresult = result
		result = ApplycaseOl(result)
	}
	return result
}
func ApplycaseOl(str string) string {
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
			result += string(runes[i : start+1])
			i = start + 1
			continue
		}
		end := Matchp(runes, start)
		if end == -1 {
			result += string(runes[i : start+1])
			i = start + 1
			continue
		}
		if end < len(runes)-1 && runes[end+1] != ' ' && runes[end+1] != '\t' && runes[end+1] != '\n' {
			result += string(runes[i : start+1])
			i = start + 1
			continue
		}
		rawcommand := string(runes[start+1 : end])
		command := strings.TrimSpace(rawcommand)
		if rawcommand != command {
			result += string(runes[i : start+1])
			i = start + 1
			continue
		}
		if strings.ContainsAny(command, "()") {
			result += string(runes[i : start+1])
			i = start + 1
			continue
		}
		var cmd string
		var num int
		validCommand := false
		if command == "hex" || command == "bin" {
			cmd = command
			num = 1
			validCommand = true
		} else if strings.Contains(command, ",") {
			if !strings.Contains(command, ", ") || strings.Contains(command, ",  ") {
				result += string(runes[i : start+1])
				i = start + 1
				continue
			}
			parts := strings.Split(command, ", ")
			if len(parts) == 2 {
				cmd = parts[0]
				numStr := parts[1]
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
		words := splitWords(allText)
		if len(words) > 0 {
			startIdx := len(words) - num
			if startIdx < 0 {
				startIdx = 0
			}
			for j := startIdx; j < len(words); j++ {
				switch cmd {
				case "low":
					words[j] = strings.ToLower(words[j])
				case "up":
					words[j] = strings.ToUpper(words[j])
				case "cap":
					words[j] = Capitalize(words[j])
				case "hex":
					words[j] = convertHex(words[j])
				case "bin":
					words[j] = convertBin(words[j])
				}
			}
			result = rebuildWithWhitespace(allText, words)
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
