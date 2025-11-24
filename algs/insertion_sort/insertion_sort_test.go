package insertion_sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{3, 7, 1}
	t.Logf("%v\n", input)
	InsertionSort(input)
	t.Logf("%v\n", input)
}
