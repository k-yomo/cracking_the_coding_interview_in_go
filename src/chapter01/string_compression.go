package main

import (
	"strconv"
	"strings"
)

// 1.6
func compressString(str string) string {
	chars := strings.Split(str, "")
	compressedStr := ""
	curChar := chars[0]
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == curChar {
			count++
		} else {
			compressedStr += curChar + strconv.Itoa(count)
			curChar = chars[i]
			count = 1
		}
	}
	compressedStr += curChar + strconv.Itoa(count)
	if len(compressedStr) >= len(str) {
		return str
	} else {
		return compressedStr
	}
}
