package ringbuffer

import (
	"sync"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	t.Run("New_EmptyPop", func(t *testing.T) {
		rb := New[int]()
		if _, ok := rb.Pop(); ok {
			t.Fatalf("Pop empty → want ok=false, got true")
		}
	})

	t.Run("NextIndex", func(t *testing.T) {
		rb := New[int]()
		if got := rb.nextIndex(19); got != 0 {
			t.Fatalf("nextIndex(19) → want 0, got %d", got)
		}
		if got := rb.nextIndex(0); got != 1 {
			t.Fatalf("nextIndex(0) → want 1, got %d", got)
		}
	})

	t.Run("PushPop_FIFO", func(t *testing.T) {
		rb := New[int]()
		for i := 0; i < 5; i++ {
			if err := rb.Push(i); err != nil {
				t.Fatalf("Push(%d) → unexpected err: %v", i, err)
			}
		}
		for i := 0; i < 5; i++ {
			v, ok := rb.Pop()
			if !ok {
				t.Fatalf("Pop #%d → want ok=true, got false", i)
			}
			if v != i {
				t.Fatalf("Pop #%d → want %d, got %d", i, i, v)
			}
		}
		if _, ok := rb.Pop(); ok {
			t.Fatalf("Pop after drain → want ok=false, got true")
		}
	})

	t.Run("Push_Full", func(t *testing.T) {
		rb := New[int]()
		for i := 0; i < 20; i++ {
			if err := rb.Push(i); err != nil {
				t.Fatalf("Push(%d) → unexpected err: %v", i, err)
			}
		}
		if err := rb.Push(99); err == nil {
			t.Fatalf("Push over cap → want err, got nil")
		} else if err.Error() != "queue already at max cap" {
			t.Fatalf("Push over cap → want %q, got %q", "queue already at max cap", err.Error())
		}
		for i := 0; i < 20; i++ {
			v, ok := rb.Pop()
			if !ok || v != i {
				t.Fatalf("Pop #%d → want (%d,true), got (%d,%v)", i, i, v, ok)
			}
		}
	})

	t.Run("Pop_EmptyReturnsZero", func(t *testing.T) {
		rb := New[string]()
		v, ok := rb.Pop()
		if ok {
			t.Fatalf("Pop empty → want ok=false, got true")
		}
		if v != "" {
			t.Fatalf("Pop empty → want zero %q, got %q", "", v)
		}
	})

	t.Run("PushFront_Basic", func(t *testing.T) {
		rb := New[int]()
		if err := rb.Push(1); err != nil {
			t.Fatal(err)
		}
		if err := rb.Push(2); err != nil {
			t.Fatal(err)
		}
		if err := rb.PushFront(0); err != nil {
			t.Fatal(err)
		}
		want := []int{0, 1, 2}
		for i, w := range want {
			v, ok := rb.Pop()
			if !ok || v != w {
				t.Fatalf("Pop #%d → want (%d,true), got (%d,%v)", i, w, v, ok)
			}
		}
	})

	t.Run("PushFront_Multiple", func(t *testing.T) {
		rb := New[int]()
		for i := 2; i >= 0; i-- {
			if err := rb.PushFront(i); err != nil {
				t.Fatalf("PushFront(%d) → unexpected err: %v", i, err)
			}
		}
		for i := 0; i < 3; i++ {
			v, ok := rb.Pop()
			if !ok || v != i {
				t.Fatalf("Pop #%d → want (%d,true), got (%d,%v)", i, i, v, ok)
			}
		}
	})

	t.Run("PushFront_Full", func(t *testing.T) {
		rb := New[int]()
		for i := 0; i < 20; i++ {
			if err := rb.PushFront(i); err != nil {
				t.Fatalf("PushFront(%d) → unexpected err: %v", i, err)
			}
		}
		if err := rb.PushFront(99); err == nil {
			t.Fatalf("PushFront over cap → want err, got nil")
		}
		if err := rb.Push(99); err == nil {
			t.Fatalf("Push over cap (after PushFront fill) → want err, got nil")
		}
	})

	t.Run("WrapAround_Head", func(t *testing.T) {
		rb := New[int]()
		for i := 0; i < 15; i++ {
			if err := rb.Push(i); err != nil {
				t.Fatal(err)
			}
		}
		for i := 0; i < 10; i++ {
			v, ok := rb.Pop()
			if !ok || v != i {
				t.Fatalf("drain #%d → want (%d,true), got (%d,%v)", i, i, v, ok)
			}
		}
		for i := 15; i < 30; i++ {
			if err := rb.Push(i); err != nil {
				t.Fatalf("Push(%d) → unexpected err: %v", i, err)
			}
		}
		for i := 10; i < 30; i++ {
			v, ok := rb.Pop()
			if !ok || v != i {
				t.Fatalf("Pop → want (%d,true), got (%d,%v)", i, v, ok)
			}
		}
	})

	t.Run("WrapAround_TailPushFront", func(t *testing.T) {
		rb := New[int]()
		if err := rb.PushFront(1); err != nil {
			t.Fatal(err)
		}
		if err := rb.PushFront(0); err != nil {
			t.Fatal(err)
		}
		for i, w := range []int{0, 1} {
			v, ok := rb.Pop()
			if !ok || v != w {
				t.Fatalf("Pop #%d → want (%d,true), got (%d,%v)", i, w, v, ok)
			}
		}
		for cycle := 0; cycle < 3; cycle++ {
			for i := 0; i < 20; i++ {
				if err := rb.PushFront(i); err != nil {
					t.Fatalf("cycle %d PushFront(%d) → %v", cycle, i, err)
				}
			}
			for i := 0; i < 20; i++ {
				v, ok := rb.Pop()
				want := 19 - i
				if !ok || v != want {
					t.Fatalf("cycle %d pop #%d → want (%d,true), got (%d,%v)", cycle, i, want, v, ok)
				}
			}
		}
	})

	t.Run("Mixed_PushPushFrontPop", func(t *testing.T) {
		rb := New[int]()
		if err := rb.Push(1); err != nil {
			t.Fatal(err)
		}
		if err := rb.Push(2); err != nil {
			t.Fatal(err)
		}
		if err := rb.PushFront(0); err != nil {
			t.Fatal(err)
		}
		v, _ := rb.Pop()
		if v != 0 {
			t.Fatalf("Pop → want 0, got %d", v)
		}
		if err := rb.Push(3); err != nil {
			t.Fatal(err)
		}
		if err := rb.PushFront(-1); err != nil {
			t.Fatal(err)
		}
		want := []int{-1, 1, 2, 3}
		for i, w := range want {
			got, ok := rb.Pop()
			if !ok || got != w {
				t.Fatalf("Pop #%d → want (%d,true), got (%d,%v)", i, w, got, ok)
			}
		}
	})

	t.Run("ReuseAfterDrain", func(t *testing.T) {
		rb := New[int]()
		for round := 0; round < 3; round++ {
			for i := 0; i < 20; i++ {
				if err := rb.Push(i + round*100); err != nil {
					t.Fatalf("round %d Push → %v", round, err)
				}
			}
			for i := 0; i < 20; i++ {
				v, ok := rb.Pop()
				if !ok || v != i+round*100 {
					t.Fatalf("round %d pop #%d → want %d, got %d (ok=%v)", round, i, i+round*100, v, ok)
				}
			}
		}
	})

	t.Run("Generic_String", func(t *testing.T) {
		rb := New[string]()
		msgs := []string{"a", "b", "c"}
		for _, m := range msgs {
			if err := rb.Push(m); err != nil {
				t.Fatal(err)
			}
		}
		for i, w := range msgs {
			v, ok := rb.Pop()
			if !ok || v != w {
				t.Fatalf("Pop #%d → want (%q,true), got (%q,%v)", i, w, v, ok)
			}
		}
	})

	t.Run("Generic_PointerCleared", func(t *testing.T) {
		rb := New[*int]()
		x, y := 1, 2
		if err := rb.Push(&x); err != nil {
			t.Fatal(err)
		}
		if err := rb.Push(&y); err != nil {
			t.Fatal(err)
		}
		v, ok := rb.Pop()
		if !ok || *v != 1 {
			t.Fatalf("Pop → want 1, got %v (ok=%v)", v, ok)
		}
		v, ok = rb.Pop()
		if !ok || *v != 2 {
			t.Fatalf("Pop → want 2, got %v (ok=%v)", v, ok)
		}
	})

	t.Run("Concurrent_PushPop", func(t *testing.T) {
		rb := New[int]()
		var wg sync.WaitGroup
		for g := 0; g < 4; g++ {
			wg.Add(1)
			go func(base int) {
				defer wg.Done()
				for i := 0; i < 5; i++ {
					_ = rb.Push(base*5 + i)
				}
			}(g)
		}
		wg.Wait()
		count := 0
		seen := map[int]bool{}
		for {
			v, ok := rb.Pop()
			if !ok {
				break
			}
			count++
			if seen[v] {
				t.Fatalf("duplicate value %d", v)
			}
			seen[v] = true
		}
		if count != 20 {
			t.Fatalf("concurrent Push → want 20 items, got %d", count)
		}
	})
}
