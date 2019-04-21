package __delete_middle_node

import "github.com/k-yomo/cracking_the_coding_interview_in_go/src/chapter02/helpers"

// 2.3 Delete Middle Node
// Implement an algorithm to delete a node in the middle
// (i.e., any node but the first and last node, not necessarily the exact middle)
// of a singly linked list, given only access to that node.
// EXAMPLE
// Input: the node c from the linked list a -> b -> c -> d -> e -> f
// Result: nothing is returned, but the new linked list looks like a -> b -> d -> e -> f
func deleteNode(n *helpers.Node) {
	if n.Next == nil {
		n = nil
		return
	}

	n.Value = n.Next.Value
	n.Next = n.Next.Next
}
