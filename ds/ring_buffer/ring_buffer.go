package ringbuffer

import (
	"errors"
	"sync"
)

type RingBuffer[T any] struct {
	head int
	tail int
	arr  []T
	size int
	m    sync.Mutex
}

func New[T any]() *RingBuffer[T] {
	return &RingBuffer[T]{arr: make([]T, 20)}
}

func (r *RingBuffer[T]) nextIndex(curr_index int) int {
	return (curr_index + 1) % len(r.arr)
}

func (r *RingBuffer[T]) Push(n T) error {
	r.m.Lock()
	defer r.m.Unlock()
	if r.size >= len(r.arr) {
		return errors.New("queue already at max cap")
	}
	r.arr[r.head] = n
	r.head = r.nextIndex(r.head)
	r.size += 1
	return nil
}

func (r *RingBuffer[T]) Pop() (T, bool) {
	var zeroT T
	r.m.Lock()
	defer r.m.Unlock()
	if r.size == 0 {
		return zeroT, false
	}
	m := r.arr[r.tail]
	// drop reference to message payload so it can be freed by GC, before we overwrrte it
	// which can happen after long time
	r.arr[r.tail] = zeroT
	r.size -= 1
	r.tail = r.nextIndex(r.tail)
	return m, true
}

func (r *RingBuffer[T]) PushFront(msg T) error {
	r.m.Lock()
	defer r.m.Unlock()
	if r.size >= len(r.arr) {
		return errors.New("queue already at max cap")
	}
	// add len of arr to properly handle negative index
	r.tail = (r.tail - 1 + len(r.arr)) % len(r.arr)
	r.arr[r.tail] = msg
	r.size++
	return nil
}
