package __one_away

import (
	"strings"
)

// 1.5 One Away
// There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character.
// Given two strings, write a function to check if they are one edit (or zero edits) away.
func isOneAway(str1 string, str2 string) bool {
	chars1, chars2 := strings.Split(str1, ""), strings.Split(str2, "")
	if len(chars1) == len(chars2) {
		return isOneEdit(chars1, chars2)
	} else if len(chars1) == len(chars2)-1 {
		return isOneInsert(chars1, chars2)
	} else if len(chars1)-1 == len(chars2) {
		return isOneInsert(chars2, chars1)
	} else {
		return false
	}
}

func isOneEdit(chars1 []string, chars2 []string) bool {
	isEdited := false
	for i := 0; i < len(chars1); i++ {
		if chars1[i] != chars2[i] {
			if isEdited {
				return false
			}
			isEdited = true
		}
	}
	return true
}

func isOneInsert(shorterChars []string, longerChars []string) bool {
	isInserted := false
	shorterCharsIndex := 0
	for i := 0; i < len(longerChars); i++ {
		if shorterCharsIndex >= len(shorterChars) {
			return true
		}
		if shorterChars[shorterCharsIndex] != longerChars[i] {
			if isInserted {
				return false
			}
			isInserted = true
		} else {
			shorterCharsIndex++
		}
	}
	return true
}
