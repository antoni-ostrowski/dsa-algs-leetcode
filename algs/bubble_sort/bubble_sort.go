package bubble_sort

// in-place sort,
// avarage/worst O(n^2)
// best (already sorted) O(n)

// good for tiny arrays <20-30

func BubbleSort(arr []int) {
	l := len(arr)
	swapped := false

	for i := 0; i < l-1; i += 1 {
		swapped = false
		for j := 0; j < l-1; j += 1 {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}

	}

}
