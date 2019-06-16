package __add_lists

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers/singly_linked_list"
	"reflect"
	"testing"
)

func TestAddLists(t *testing.T) {
	testCases := []struct {
		Node1    *singly_linked_list.Node
		Node2    *singly_linked_list.Node
		Expected *singly_linked_list.SinglyLinkedList
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			singly_linked_list.NewFromArray([]interface{}{1, 9, 4}).Front(),
			singly_linked_list.NewFromArray([]interface{}{3, 1, 9}).Front(),
			singly_linked_list.NewFromArray([]interface{}{4, 0, 3}),
		},
	}

	for _, tc := range testCases {
		result := addLists(tc.Node1, tc.Node2)
		if !reflect.DeepEqual(tc.Expected, result) {
			t.Fatalf("failed test expected: %v, got: %v", tc.Expected, result)
		}
	}
}
