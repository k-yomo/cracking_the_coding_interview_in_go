package main

import (
	"fmt"
	"strings"
)

// 1.3
func main()  {
	fmt.Println(urlify("Mr John Smith   ", 13))
}

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
