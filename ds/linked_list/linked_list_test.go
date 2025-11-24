package linked_list

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("InsertHead", func(t *testing.T) {
		var l LinkedList
		l.InsertHead(3)
		l.InsertHead(2)
		l.InsertHead(1)

		if got := l.GetValues(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
			t.Fatalf("GetValues = %v, want [1 2 3]", got)
		}
	})

	t.Run("InsertTail", func(t *testing.T) {
		var l LinkedList
		l.InsertTail(1)
		l.InsertTail(2)
		l.InsertTail(3)

		if got := l.GetValues(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
			t.Fatalf("GetValues = %v, want [1 2 3]", got)
		}
	})

	t.Run("Get", func(t *testing.T) {
		var l LinkedList
		l.InsertTail(10)
		l.InsertTail(20)
		l.InsertTail(30)

		if got := l.Get(0); got != 10 {
			t.Fatalf("Get(0) = %d, want 10", got)
		}
		if got := l.Get(2); got != 30 {
			t.Fatalf("Get(2) = %d, want 30", got)
		}
		if got := l.Get(-1); got != -1 {
			t.Fatalf("Get(-1) = %d, want -1", got)
		}
		if got := l.Get(3); got != -1 {
			t.Fatalf("Get(3) = %d, want -1", got)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		var l LinkedList
		l.InsertTail(1)
		l.InsertTail(2)
		l.InsertTail(3)

		if !l.Remove(0) {
			t.Fatal("Remove(0) = false, want true")
		}
		if got := l.GetValues(); !reflect.DeepEqual(got, []int{2, 3}) {
			t.Fatalf("after Remove(0) = %v, want [2 3]", got)
		}

		if !l.Remove(1) {
			t.Fatal("Remove(1) = false, want true")
		}
		if got := l.GetValues(); !reflect.DeepEqual(got, []int{2}) {
			t.Fatalf("after Remove(1) = %v, want [2]", got)
		}

		if l.Remove(5) {
			t.Fatal("Remove(5) = true, want false")
		}
		if l.Remove(-1) {
			t.Fatal("Remove(-1) = true, want false")
		}
	})

	t.Run("GetValuesEmpty", func(t *testing.T) {
		var l LinkedList
		got := l.GetValues()
		if got == nil {
			t.Fatal("GetValues = nil, want empty non-nil slice")
		}
		if len(got) != 0 {
			t.Fatalf("len(GetValues) = %d, want 0", len(got))
		}
	})
}
