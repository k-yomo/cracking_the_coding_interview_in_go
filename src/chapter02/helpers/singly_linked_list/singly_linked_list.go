package singly_linked_list

import (
	"fmt"
	"strings"
)

type SinglyLinkedList struct {
	root *Node
	len  int
}

func (s *SinglyLinkedList) Init() *SinglyLinkedList {
	s.root = nil
	s.len = 0
	return s
}

func New() *SinglyLinkedList { return new(SinglyLinkedList).Init() }

func NewFromArray(array []interface{}) *SinglyLinkedList {
	list := New()
	for _, value := range array {
		list.PushBack(&Node{Value: value})
	}
	return list
}

func (s *SinglyLinkedList) Len() int {
	return s.len
}

func (s *SinglyLinkedList) Front() *Node {
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
		s.len = 1
	} else {
		n := s.root

		for n.Next != nil {
			n = n.Next
		}

		n.Next = appendedNode
		s.len++
	}
}

func (s *SinglyLinkedList) ToArray() []interface{} {
	if s.len == 0 {
		return []interface{}{}
	}
	array := make([]interface{}, 0, s.len)
	n := s.root
	for n != nil {
		array = append(array, n.Value)
		n = n.Next
	}
	return array
}

func (s *SinglyLinkedList) String() string {
	array := s.ToArray()
	strArray := make([]string, len(array))
	for i, item := range array {
		strArray[i] = fmt.Sprintf("%v", item)
	}
	return fmt.Sprintf("SinglyLinkedList{%s}", strings.Join(strArray, "->"))
}
