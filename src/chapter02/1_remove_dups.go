package chapter02

import (
	"container/list"
)

// 2.1 Remove Dups
// Write code to remove duplicates from an unsorted linked list
// How would you solve this problem if a temporary buffer is not allowed?
// space complexity: O(1), time complexity: O(n²)
func removeDupsFromList(l *list.List) {
	for el := l.Front(); el != nil; el = el.Next() {
		elExist := false
		for innerEl := l.Front(); innerEl != nil; innerEl = innerEl.Next() {
			if el.Value != innerEl.Value {
				continue
			} else if elExist == true {
				l.Remove(innerEl)
			} else {
				elExist = true
			}
		}
	}
}

// space complexity: O(n), time complexity: O(n)
func removeDupsFromListWithBuffer(l *list.List) {
	elMap := make(map[interface{}]bool, 0)
	for el := l.Front(); el != nil; el = el.Next() {
		if elMap[el.Value] {
			l.Remove(el)
		} else {
			elMap[el.Value] = true
		}
	}
}
