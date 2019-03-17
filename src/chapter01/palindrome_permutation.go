package main

import (
	"strings"
)

// 1.4
func checkPalindromePermutation(str string) bool  {
	chars := strings.Split(str, "")
	charMap := make(map[string]int)
	for _, char := range chars  {
		charMap[char] += 1
	}
	nonUniqueChars := 0
	for _, value := range charMap {
		if value % 2 == 1 {
			nonUniqueChars++
		}
	}
	return nonUniqueChars <= 2
}

