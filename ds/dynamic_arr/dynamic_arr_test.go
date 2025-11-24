package dynamic_arr

import "testing"

func TestDynamicArr(t *testing.T) {
	t.Run("GrowOrder", func(t *testing.T) {
		d := NewDynamicArray(2)
		for i := range 10 {
			d.PushBack(i)
		}
		for i := range 10 {
			if v, _ := d.Get(i); v != i {
				t.Fatalf("Get(%d)=%d want %d", i, v, i)
			}
		}
	})

	t.Run("GetOOR", func(t *testing.T) {
		d := NewDynamicArray(2)
		for _, i := range []int{-1, 0, 5} {
			if _, ok := d.Get(i); ok {
				t.Fatalf("Get(%d) ok=true want false", i)
			}
		}
	})

	t.Run("Set", func(t *testing.T) {
		d := NewDynamicArray(2)
		d.PushBack(1)
		if !d.Set(0, 9) {
			t.Fatal("Set(0)=false want true")
		}
		if v, _ := d.Get(0); v != 9 {
			t.Fatalf("Get(0)=%d want 9", v)
		}
		if d.Set(1, 9) || d.Set(-1, 9) {
			t.Fatal("Set OOR=true want false")
		}
	})

	t.Run("Pop", func(t *testing.T) {
		d := NewDynamicArray(2)
		if _, ok := d.PopBack(); ok {
			t.Fatal("Pop empty=true want false")
		}
		d.PushBack(1)
		d.PushBack(2)
		if v, _ := d.PopBack(); v != 2 {
			t.Fatalf("Pop=%d want 2", v)
		}
		d.PushBack(3)
		if v, _ := d.Get(1); v != 3 {
			t.Fatalf("Get(1)=%d want 3", v)
		}
	})
}
