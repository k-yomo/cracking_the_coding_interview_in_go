package __add_lists

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers"
)

// 2.4 Add Lists
// You have two numbers represented by a linked list, where each node contains a single digit.
// The digits are stored in reverse order, such that the 1's digit is at the head of the list.
// Write a function that adds the two numbers and returns the sum as a linked list.
func addLists(node1 *helpers.Node, node2 *helpers.Node) *helpers.SinglyLinkedList {
	if node1 == nil && node2 == nil {
		return nil
	}

	linkedList := helpers.NewSinglyLinkedList()
	for node1 != nil {
		linkedList.PushBack(&helpers.Node{Value: (node1.Value.(int) + node2.Value.(int)) % 10})
		node1 = node1.Next
		node2 = node2.Next
	}
	return linkedList
}
