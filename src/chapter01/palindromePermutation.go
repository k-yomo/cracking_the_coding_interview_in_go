package main

import (
	"fmt"
	"strings"
)

// 1.4
func main() {
	fmt.Println(checkPalindromePermutation("taco cat"))
	fmt.Println(checkPalindromePermutation("taco cats"))
}

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

