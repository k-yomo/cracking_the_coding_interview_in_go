package __delete_middle_node

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers/singly_linked_list"
	"reflect"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	singlyLinkedList := singly_linked_list.New()
	node1 := &singly_linked_list.Node{Value: "a"}
	node2 := &singly_linked_list.Node{Value: "b"}
	node3 := &singly_linked_list.Node{Value: "c"}
	singlyLinkedList.PushBack(node1)
	singlyLinkedList.PushBack(node2)
	singlyLinkedList.PushBack(node3)

	testCases := []struct {
		priorNode *singly_linked_list.Node
		node      *singly_linked_list.Node
		expected  *singly_linked_list.Node
	}{
		{
			node1,
			node2,
			node3,
		},
		{
			node2,
			node3,
			nil,
		},
	}
	for _, tc := range testCases {
		deleteNode(tc.node)
		if !reflect.DeepEqual(tc.expected, tc.priorNode.Next) {
			t.Fatalf("failed test expected: %v, got: %v", tc.expected, tc.priorNode.Next)
		}
	}
}
