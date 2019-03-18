package __string_compression

import (
	"strconv"
	"strings"
)

// 1.6 String Compression
// Implement a method to perform basic string compression using the counts of repeated characters.
// For example, the string aabcccccaaa would become a2b1c5a3.
// If the "compressed" string would not become smaller than the original string, your method should return the original string.
// You can assume the string has only uppercase and lowercase letters (a-z) .
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
