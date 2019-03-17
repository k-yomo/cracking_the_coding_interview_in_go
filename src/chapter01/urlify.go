package main

import (
	"strings"
)

// 1.3
func urlify(str string, charLength int) string {
	var chars []string
	for i, charRune := range str {
		if i >= charLength {
			break
		}
		char := string(charRune)
		if char == " " {
			chars = append(chars, "%20")
		} else {
			chars = append(chars, char)
		}
	}
	return strings.Join(chars, "")
}
