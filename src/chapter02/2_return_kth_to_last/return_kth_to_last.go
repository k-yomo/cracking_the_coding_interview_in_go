package __return_kth_to_last

import (
	"github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers/singly_linked_list"
)

// 2.2 Return Kth to Last
// Implement an algorithm to find the kth to last element of a singly linked list.
// We assume kth means kth index that counts from 0.
func findKthToLast(list *singly_linked_list.SinglyLinkedList, kth int) []*singly_linked_list.Node {
	curNode := list.Front()
	nodes := make([]*singly_linked_list.Node, 0)
	i := 0
	for curNode != nil {
		if i >= kth {
			nodes = append(nodes, curNode)
		}
		curNode = curNode.Next
		i++
	}
	return nodes
}
