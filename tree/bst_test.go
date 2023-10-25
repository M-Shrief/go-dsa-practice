package tree

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	t.Run("Test Insert()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)

		want := int(5)
		n := tree.GetRoot()
		got := n.GetVal()
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want1 := int(2)
		n1 := n.left
		got1 := n1.GetVal()
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}

		want2 := int(1)
		n2 := n1.left
		got2 := n2.GetVal()
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(3)
		n3 := n1.right
		got3 := n3.GetVal()
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want4 := int(7)
		n4 := n.right
		got4 := n4.GetVal()
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got4, want4)
		}

		want5 := int(6)
		n5 := n4.left
		got5 := n5.GetVal()
		if got5 != want5 {
			t.Errorf("Got: %v, Want: %v", got5, want5)
		}

		want6 := int(8)
		n6 := n4.right
		got6 := n6.GetVal()
		if got6 != want6 {
			t.Errorf("Got: %v, Want: %v", got6, want6)
		}
	})

	t.Run("Test Search()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(6)

		want := int(5)
		n, ok := tree.Search(5)
		got := n.GetVal()

		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want1 := int(2)
		n1, ok1 := tree.Search(2)
		got1 := n1.GetVal()

		if !ok1 {
			t.Error("Not Okay")
		}
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}

		want2 := int(1)
		n2, ok2 := tree.Search(1)
		got2 := n2.GetVal()
		if !ok2 {
			t.Error("Not Okay")
		}
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(3)
		n3, ok3 := tree.Search(3)
		got3 := n3.GetVal()
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

		want4 := int(6)
		n4, ok4 := tree.Search(6)
		got4 := n4.GetVal()
		if !ok4 {
			t.Error("Not Okay")
		}
		if got4 != want4 {
			t.Errorf("Got: %v, Want: %v", got4, want4)
		}
	})
	t.Run("Test minimunVal()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want := int(1)
		got := tree.Minimum()

		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test minimunVal()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want := int(8)
		got := tree.Maximum()

		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
	t.Run("Test GetParent()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want := tree.root
		got := tree.GetParent(2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want1 := tree.root.left
		got1 := tree.GetParent(1)
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}
		got2 := tree.GetParent(3)
		if !reflect.DeepEqual(got2, want1) {
			t.Errorf("Got: %v, Want: %v", got2, want1)
		}

		want2 := tree.root.right
		got3 := tree.GetParent(6)
		if !reflect.DeepEqual(got3, want2) {
			t.Errorf("Got: %v, Want: %v", got3, want2)
		}
		got4 := tree.GetParent(8)
		if !reflect.DeepEqual(got4, want2) {
			t.Errorf("Got: %v, Want: %v", got4, want2)
		}
	})

}
