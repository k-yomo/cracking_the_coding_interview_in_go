package __is_unique

import (
	"sort"
	"strings"
)

// 1.1 Is Unique
// Implement an algorithm to determine if a string has all unique characters.
// What is you cannot use additional data structures?
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
	for _, char := range []rune(str) {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}
