package helpers

type SinglyLinkedList struct {
	root *Node
	len  int
}

func (s *SinglyLinkedList) Init() *SinglyLinkedList {
	s.root = nil
	s.len = 0
	return s
}

func NewSinglyLinkedList() *SinglyLinkedList { return new(SinglyLinkedList).Init() }

func (s *SinglyLinkedList) Len() int {
	return s.len
}

func (s *SinglyLinkedList) Front() *Node {
	if s.len == 0 {
		return nil
	}
	return s.root
}

func (s *SinglyLinkedList) Back() *Node {
	if s.len == 0 {
		return nil
	}
	n := s.root
	for n.Next != nil {
		n = n.Next
	}
	return n
}

func (s *SinglyLinkedList) PushBack(appendedNode *Node) {
	if s.root == nil {
		s.root = appendedNode
	} else {
		n := s.root

		for n.Next != nil {
			n = n.Next
		}

		n.Next = appendedNode
	}
	s.len++
}
