package __urlify

import (
	"strings"
)

// 1.3 URLify
// Write a method to replace all spaces in a string with '%20'.
// You may assume that the string has sufficient space at the end to hold the additional characters,
// and that you are given the "true" length of the string.
func urlify(str string, charLength int) string {
	chars := make([]string, charLength)
	for i, charRune := range str {
		if i >= charLength {
			break
		}
		char := string(charRune)
		if char == " " {
			chars[i] = "%20"
		} else {
			chars[i] = char
		}
	}
	return strings.Join(chars, "")
}
