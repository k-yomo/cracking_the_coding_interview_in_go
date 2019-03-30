package __remove_dups

import (
	"container/list"
	"reflect"
	"testing"
)

func TestRemoveDupsFromList(t *testing.T) {
	dupList := list.New()
	dupList.PushBack("a")
	dupList.PushBack("b")
	dupList.PushBack("a")
	dupList.PushBack("c")
	expectedList := list.New()
	expectedList.PushBack("a")
	expectedList.PushBack("b")
	expectedList.PushBack("c")

	testCases := []struct {
		list     *list.List
		expected *list.List
	}{
		{
			dupList,
			expectedList,
		},
	}
	for _, tc := range testCases {
		removeDupsFromList(tc.list)
		if !reflect.DeepEqual(tc.expected, tc.list) {
			t.Fatalf("failed test expected: %v, got: %v", tc.expected, tc.list)
		}
	}
}
func TestRemoveDupsFromListWithBuffer(t *testing.T) {
	dupList := list.New()
	dupList.PushBack("a")
	dupList.PushBack("b")
	dupList.PushBack("a")
	dupList.PushBack("c")
	expectedList := list.New()
	expectedList.PushBack("a")
	expectedList.PushBack("b")
	expectedList.PushBack("c")

	testCases := []struct {
		list     *list.List
		expected *list.List
	}{
		{
			dupList,
			expectedList,
		},
	}
	for _, tc := range testCases {
		removeDupsFromList(tc.list)
		if !reflect.DeepEqual(tc.expected, tc.list) {
			t.Fatalf("failed test expected: %v, got: %v", tc.expected, tc.list)
		}
	}
}
