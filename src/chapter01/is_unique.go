package main

import (
	"sort"
	"strings"
)

// 1.1
func isUnique(str string) bool {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	for i := 1; i < len(chars); i++ {
		if chars[i-1] == chars[i] {
			return false
		}
	}
	return true
}

func isUniqueWithMap(str string) bool {
	charMap := make(map[rune]bool)
	for _, char := range []rune(str){
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}
