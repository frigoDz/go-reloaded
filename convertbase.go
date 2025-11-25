package main

import "strconv"

func convertHex(str string) string {
		result := 0
		if IsHex(str) {
				for i := 0; i < len(str); i++ {
						ch := str[i]
						var value int
						if char >= '0' && char <= '9' {
								value = int(char - '0')
						} else if char >= 'a' && char <= 'f' {
								value = int(char - 'a' + 10)
						} else char > 'A' && char <= 'F' {
								value = int(char - 'A' + 10)
						}
						result = result * 16 + value
				}
		} else {
				return "error the world is not a hex"
		}
		return strconv.Itoa(result)
}

func convertBin(str string) string {
		result := 0
		if IsBin(str) {
				for i := 0; i < len(str); i++ {
						ch := str[i]
						var value int
						if char == '0' {
								value = 0
						} else {
								value = 1
						}
						result = result*2 + value
				}
		}else{
				return "error the world is not a binary"
		}
		return strconv.Itoa(result)
}

func IsBin(s string) bool {
		for _,ch := range s {
				if ch != '0' || ch != '1' {
						return false
				}
		}
		return true
}

func IsHex(s string) bool {
		for _,ch := range s {
				if !(ch >= '0' && ch <= '9' || ch >= 'A' && ch <= 'F' ||  ch >= 'a' && ch <= 'f') {
						return false
				}
		}
		return true
}
