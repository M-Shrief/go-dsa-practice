package linkedlist

import (
	"fmt"
)

type SinglyNode[T any] struct {
	val  T
	next *SinglyNode[T]
}

func NewSinglyNode[T any](val T) *SinglyNode[T] {
	return &SinglyNode[T]{val, nil}
}

func (n *SinglyNode[T]) GetVal() T {
	return n.val
}

func (n *SinglyNode[T]) GetNext() *SinglyNode[T] {
	return n.next
}

type Singly[T any] struct {
	size int
	head *SinglyNode[T]
	tail *SinglyNode[T]
}

func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

func (list *Singly[T]) GetHead() *SinglyNode[T] {
	return list.head
}

func (list *Singly[T]) GetTail() *SinglyNode[T] {
	if list.size == 0 {
		return nil
	}
	return list.tail
}

func (list *Singly[T]) GetSize() int {
	return list.size
}

func (list *Singly[T]) GetNode(pos int) any {

	if pos < 0 || pos >= list.size {
		return nil
	}
	if pos == 0 {
		return list.head.val
	}
	if pos == list.size-1 {
		return list.tail.val
	}

	current := list.head

	for count := 0; count < pos; count++ {
		current = current.next
	}

	return current.val
}

func (list *Singly[T]) AddFirst(val T) {
	n := NewSinglyNode(val)
	if list.size == 0 {
		list.head = n
		list.tail = n
	} else {
		n.next = list.head
		list.head = n
	}
	list.size++
}

func (list *Singly[T]) AddLast(val T) {
	n := NewSinglyNode(val)

	if list.size == 0 {
		list.AddFirst(val)
		return
	} else {
		list.tail.next = n
		list.tail = n
	}

	list.size++
}

func (list *Singly[T]) DeleteFirst() (T, bool) {
	if list.head == nil {
		var r T
		return r, false
	}
	current := list.head
	if list.size == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.head = current.next
	}

	list.size--
	return current.val, true
}

func (list *Singly[T]) DeleteLast() (T, bool) {
	return list.DeleteNode(list.size - 1)
}

func (list *Singly[T]) DeleteNode(pos int) (T, bool) {

	if pos < 0 || pos >= list.size {
		var r T
		return r, false
	}
	if pos == 0 {
		return list.DeleteFirst()
	}

	current := list.head

	for count := 0; count < pos-1; count++ {
		current = current.next
	}

	removed := current.next.val
	current.next = current.next.next
	list.size--

	return removed, true
}

func (list *Singly[T]) Reverse() {
	var prev, next *SinglyNode[T]
	current := list.head

	for current != nil {
		next = current.next
		current.next = prev

		prev = current
		current = next
	}

	list.head = prev
}

func (list *Singly[T]) Display() {
	for current := list.head; current != nil; current = current.next {
		fmt.Println(current.val)
	}
}
