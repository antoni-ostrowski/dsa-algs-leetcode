package linked_list

type node struct {
	val  int
	next *node
}
type LinkedList struct {
	head *node
	size int
}

func (l *LinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	count := 0
	curr := l.head
	for count < index {
		curr = curr.next
		count++
	}
	return curr.val
}

func (l *LinkedList) InsertHead(val int) {
	newNode := &node{val: val}
	newNode.next = l.head
	l.head = newNode
	l.size++
}

func (l *LinkedList) InsertTail(val int) {
	newTailNode := &node{val: val}
	if l.head == nil {
		l.head = newTailNode
		l.size++
		return
	}
	curr := l.head
	// finding tail node
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newTailNode
	l.size++
}

func (l *LinkedList) Remove(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	if index == 0 {
		l.head = l.head.next
		l.size--
		return true
	}
	count := 0
	curr := l.head
	// walking list until we hit desired index
	for count < index-1 {
		curr = curr.next
		count++
	}
	curr.next = curr.next.next
	l.size--
	return true
}
func (l *LinkedList) GetValues() []int {
	values := make([]int, 0, l.size)
	curr := l.head
	for curr != nil {
		values = append(values, curr.val)
		curr = curr.next
	}
	return values
}
