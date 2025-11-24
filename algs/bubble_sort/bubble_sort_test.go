package bubble_sort

import "testing"

func TestBubbleSort(t *testing.T) {
	input := []int{3, 7, 1, 10, 20, 1, 4}
	t.Logf("%v\n", input)
	BubbleSort(input)
	t.Logf("%v\n", input)
}
