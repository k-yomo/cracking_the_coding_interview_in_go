package helpers

type SinglyLinkedList struct {
	root *Node
	len int
}

func (s *SinglyLinkedList) Init() *SinglyLinkedList {
	s.root = nil
	s.len = 0
	return s
}

func New() *SinglyLinkedList { return new(SinglyLinkedList).Init() }

func (s *SinglyLinkedList) Len() int {
	return s.len
}

func (s *SinglyLinkedList) Front() *Node {
	if s.len == 0 {
		return nil
	}
	return s.root.next
}

func (s *SinglyLinkedList) Back() *Node {
	if s.len == 0 {
		return nil
	}
	n := s.root
	for n.next != nil {
		n = n.next
	}
	return n
}

func (s *SinglyLinkedList) Append(appendedNode *Node) {
	if s.root == nil {
		s.root = appendedNode
	} else {
		n := s.root

		for n.next != nil {
			n = n.next
		}

		n.next = appendedNode
	}
	s.len++
}


