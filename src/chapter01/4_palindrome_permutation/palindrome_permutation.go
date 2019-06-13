package __palindrome_permutation

import (
	"strings"
)

// 1.4 Palindrome Permutation
// Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
func checkPalindromePermutation(str string) bool {
	chars := strings.Split(str, "")
	charMap := make(map[string]int)
	for _, char := range chars {
		charMap[char] += 1
	}
	nonUniqueChars := 0
	for _, value := range charMap {
		if value%2 == 1 {
			nonUniqueChars++
		}
	}
	return nonUniqueChars <= 2
}
