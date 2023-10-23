package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoubly(t *testing.T) {
	t.Run("Test AddFirst()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := int(2)
		got := list.GetNode(1).(*DoublyNode[int])
		if got.Val != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		n1 := list.GetHead()
		val1 := n1.Val
		want1 := int(3)
		if val1 != want1 {
			t.Errorf("Got: %v, Want: %v", val1, want1)
		}
		n2 := n1.Next
		val2 := n2.Val
		want2 := int(2)
		if val2 != want2 {
			t.Errorf("Got: %v, Want: %v", val2, want2)
		}
		want3 := list.GetHead()
		val3 := n2.Prev
		if !reflect.DeepEqual(val3, want3) {
			t.Errorf("Got: %v, Want: %v", val3, want3)
		}
		want4 := list.GetTail()
		val4 := n2.Next
		if !reflect.DeepEqual(val4, want4) {
			t.Errorf("Got: %v, Want: %v", val4, want4)
		}
	})
	t.Run("Test AddLast()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)

		want := int(2)
		got := list.GetNode(1).(*DoublyNode[int]).Val
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		n1 := list.GetHead()
		val1 := n1.Val
		want1 := int(1)
		if val1 != want1 {
			t.Errorf("Got: %v, Want: %v", val1, want1)
		}
		n2 := n1.Next
		val2 := n2.Val
		want2 := int(2)
		if val2 != want2 {
			t.Errorf("Got: %v, Want: %v", val2, want2)
		}
		want3 := list.GetHead()
		val3 := n2.Prev
		if !reflect.DeepEqual(val3, want3) {
			t.Errorf("Got: %v, Want: %v", val3, want3)
		}
		want4 := list.GetTail()
		val4 := n2.Next
		if !reflect.DeepEqual(val4, want4) {
			t.Errorf("Got: %v, Want: %v", val4, want4)
		}
	})
	t.Run("Test GetNode()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddFirst(4)

		want1 := list.GetHead()
		val1 := list.GetNode(0)
		if !reflect.DeepEqual(val1, want1) {
			t.Errorf("Got: %v, Want: %v", val1, want1)
		}

		want2 := want1.Next
		val2 := list.GetNode(1)
		if !reflect.DeepEqual(val2, want2) {
			t.Errorf("Got: %v, Want: %v", val2, want2)
		}

		want3 := want2.Next
		val3 := list.GetNode(2)
		if !reflect.DeepEqual(val3, want3) {
			t.Errorf("Got: %v, Want: %v", val3, want3)
		}

		want4 := list.GetTail()
		val4 := list.GetNode(3)
		if !reflect.DeepEqual(val4, want4) {
			t.Errorf("Got: %v, Want: %v", val4, want4)
		}
	})
}
