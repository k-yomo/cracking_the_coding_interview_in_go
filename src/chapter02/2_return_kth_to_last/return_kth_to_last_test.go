package __return_kth_to_last

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers"
	"reflect"
	"testing"
)

func TestFindKthToLast(t *testing.T) {
	singlyLinkedList := helpers.NewSinglyLinkedList()
	node1 := &helpers.Node{Value: "a"}
	node2 := &helpers.Node{Value: "b"}
	node3 := &helpers.Node{Value: "c"}
	singlyLinkedList.PushBack(node1)
	singlyLinkedList.PushBack(node2)
	singlyLinkedList.PushBack(node3)
	expectedNodes := []*helpers.Node{node2, node3}

	testCases := []struct {
		list     *helpers.SinglyLinkedList
		kth int
		expected []*helpers.Node
	}{
		{
			singlyLinkedList,
			1,
			expectedNodes,
		},
	}
	for _, tc := range testCases {
		result := findKthToLast(tc.list, tc.kth)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v", tc.expected, result)
		}
	}
}
