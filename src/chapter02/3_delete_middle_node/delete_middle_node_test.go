package __delete_middle_node

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers"
	"reflect"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	singlyLinkedList := helpers.NewSinglyLinkedList()
	node1 := &helpers.Node{Value: "a"}
	node2 := &helpers.Node{Value: "b"}
	node3 := &helpers.Node{Value: "c"}
	singlyLinkedList.PushBack(node1)
	singlyLinkedList.PushBack(node2)
	singlyLinkedList.PushBack(node3)

	testCases := []struct {
		priorNode *helpers.Node
		node     *helpers.Node
		expected *helpers.Node
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
