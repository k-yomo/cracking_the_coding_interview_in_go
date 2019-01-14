package main

import (
	"fmt"
	"sort"
	"strings"
)

// 1.2
func main() {
	fmt.Println(checkPermutation("abc", "cba"))
	fmt.Println(checkPermutation("abca", "cba"))
	fmt.Println(checkPermutation("abcb", "cba"))
}

func checkPermutation( str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	chars1 := strings.Split(str1, "")
	chars2 := strings.Split(str2, "")
	sort.Strings(chars1)
	sort.Strings(chars2)
	for i := 0; i < len(chars1); i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}
	return true
}