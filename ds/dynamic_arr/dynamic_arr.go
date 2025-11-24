package dynamic_arr

type DynamicArr struct {
	capacity int
	length   int
	arr      []int
}

func NewDynamicArray(nc int) *DynamicArr {
	if nc < 0 {
		nc = 0
	}
	return &DynamicArr{capacity: nc, length: 0, arr: make([]int, nc)}
}
func (d *DynamicArr) GetLength() int {
	return d.length
}
func (d *DynamicArr) GetCap() int {
	return d.capacity
}

func (d *DynamicArr) Get(i int) (int, bool) {
	if i >= d.length || i < 0 {
		return 0, false
	}

	return d.arr[i], true
}
func (d *DynamicArr) Set(i int, v int) bool {
	if i >= d.length || i < 0 {
		return false
	}
	d.arr[i] = v
	return true
}
func (d *DynamicArr) PushBack(n int) {
	if d.capacity == d.length {
		nc := d.capacity * 2
		if nc == 0 {
			nc = 1
		}
		newArr := make([]int, nc)
		copy(newArr, d.arr)
		d.arr = newArr
		d.capacity = nc
	}
	d.arr[d.length] = n
	d.length += 1
}
func (d *DynamicArr) PopBack() (int, bool) {
	if d.length < 1 {
		return 0, false
	}
	d.length -= 1
	return d.arr[d.length], true
}
