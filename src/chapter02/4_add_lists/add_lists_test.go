package __add_lists

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers"
	"reflect"
	"testing"
)

func TestAddLists(t *testing.T) {
	testCases := []struct {
		Node1    *helpers.Node
		Node2    *helpers.Node
		Expected *helpers.SinglyLinkedList
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			helpers.NewSinglyLinkedListFromIntArray([]int{1, 9, 4}).Front(),
			helpers.NewSinglyLinkedListFromIntArray([]int{3, 1, 9}).Front(),
			helpers.NewSinglyLinkedListFromIntArray([]int{4, 0, 3}),
		},
	}

	for _, tc := range testCases {
		result := addLists(tc.Node1, tc.Node2)
		if !reflect.DeepEqual(tc.Expected, result) {
			t.Fatalf("failed test expected: %v, got: %v", tc.Expected, result)
		}
	}
}
