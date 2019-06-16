package __return_kth_to_last

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers/singly_linked_list"
	"reflect"
	"testing"
)

func TestFindKthToLast(t *testing.T) {
	singlyLinkedList := singly_linked_list.New()
	node1 := &singly_linked_list.Node{Value: "a"}
	node2 := &singly_linked_list.Node{Value: "b"}
	node3 := &singly_linked_list.Node{Value: "c"}
	singlyLinkedList.PushBack(node1)
	singlyLinkedList.PushBack(node2)
	singlyLinkedList.PushBack(node3)
	expectedNodes := []*singly_linked_list.Node{node2, node3}

	testCases := []struct {
		list     *singly_linked_list.SinglyLinkedList
		kth      int
		expected []*singly_linked_list.Node
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
