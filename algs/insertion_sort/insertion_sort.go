package insertion_sort

// https://www.youtube.com/watch?v=IniSptbttgg

// in-place sort,
// avarage/worst O(n^2)
// best (already sorted) O(n)

// good for tiny arrays <20-30

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i += 1 {
		key := arr[i]

		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
}
