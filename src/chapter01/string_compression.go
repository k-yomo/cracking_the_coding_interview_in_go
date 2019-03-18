package main

import (
	"strconv"
	"strings"
)

func main() {
	compressString("aaaabbcddda")
}

// 1.6
// return original string if original string is shorter than compressed string
func compressString(str string) string {
	chars := strings.Split(str, "")
	var compressedStr strings.Builder
	consecutiveCount := 0

	for i := 0; i < len(chars); i++ {
		consecutiveCount++
		if i+1 >= len(chars) || chars[i] != chars[i+1] {
			compressedStr.WriteString(chars[i])
			compressedStr.WriteString(strconv.Itoa(consecutiveCount))
			consecutiveCount = 0
		}
	}
	if len(compressedStr.String()) >= len(str) {
		return str
	} else {
		return compressedStr.String()
	}
}
