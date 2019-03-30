package helpers

type Node struct {
	Value interface{}
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}
