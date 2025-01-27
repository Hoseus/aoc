package utils

import (
	"cmp"
	"container/heap"
	"container/list"
	"strconv"
)

type Tuple[L, R any] struct {
	Left  L
	Right R
}

type Heap[T any] struct {
	data   []T
	isLess func(T, T) bool
}

func NewHeap[T any](isLess func(T, T) bool) *Heap[T] {
	h := &Heap[T]{
		isLess: isLess,
		data:   make([]T, 0),
	}
	heap.Init(h)
	return h
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Less(i, j int) bool {
	return h.isLess(h.data[i], h.data[j])
}

func (h *Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) PushT(x T) {
	heap.Push(h, x)
}

func (h *Heap[T]) Push(x any) {
	h.data = append(h.data, x.(T))
}

func (h *Heap[T]) PopT() T {
	return heap.Pop(h).(T)
}

func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *Heap[T]) PeekT() T {
	return h.data[0]
}

func (h *Heap[T]) Clone() Heap[T] {
	newHeap := Heap[T]{data: make([]T, len(h.data)), isLess: h.isLess}
	copy(newHeap.data, h.data) // Deep copy of the slice
	return newHeap
}

type Deque[T any] struct {
	list *list.List
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{list: list.New()}
}

func (d *Deque[T]) PushFront(value T) {
	d.list.PushFront(value)
}

func (d *Deque[T]) PushBack(value T) {
	d.list.PushBack(value)
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.list.Len() == 0 {
		var zero T
		return zero, false
	}
	front := d.list.Front()
	d.list.Remove(front)
	return front.Value.(T), true
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.list.Len() == 0 {
		var zero T
		return zero, false
	}
	back := d.list.Back()
	d.list.Remove(back)
	return back.Value.(T), true
}

func (d *Deque[T]) Len() int {
	return d.list.Len()
}

func Abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func Even(num int) bool {
	return num%2 == 0
}

func StringsToNumbers(strings []string) ([]int, error) {
	numbers := make([]int, len(strings))
	for i, aString := range strings {
		number, err := strconv.Atoi(aString)
		if err != nil {
			return nil, err
		}
		numbers[i] = number
	}
	return numbers, nil
}

func QuickSort[T cmp.Ordered](arr *[]T, isLess func(T, T) bool) {
	quickSortHelper(arr, 0, len(*arr)-1, isLess)
}

func quickSortHelper[T cmp.Ordered](arr *[]T, begin int, end int, isLess func(T, T) bool) {
	if begin < end {
		partitionIndex := quickSortPartition(arr, begin, end, isLess)

		quickSortHelper(arr, begin, partitionIndex-1, isLess)
		quickSortHelper(arr, partitionIndex+1, end, isLess)
	}
}

func quickSortPartition[T cmp.Ordered](arr *[]T, begin int, end int, isLess func(T, T) bool) int {
	pivot := (*arr)[end]
	i := begin - 1

	for j := begin; j < end; j++ {
		if isLess((*arr)[j], pivot) {
			i++

			swapTemp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = swapTemp
		}
	}

	swapTemp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[end]
	(*arr)[end] = swapTemp

	return i + 1
}

func BinarySearch[T cmp.Ordered](arr *[]T, begin int, end int, x T) int {
	for begin <= end {
		mid := begin + (end-begin)/2

		if (*arr)[mid] == x {
			return mid
		}

		if (*arr)[mid] < x {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
