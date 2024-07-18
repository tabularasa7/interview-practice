package heap

import "fmt"

type MaxHeap struct {
	values []int
}

type MinHeap struct {
	values []int
}

func NewMaxHeap(values []int) *MaxHeap {
	h := &MaxHeap{}
	h.values = values
	for i := len(values) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}

	return h
}

func NewMinHeap(values []int) *MinHeap {
	h := &MinHeap{}
	h.values = values
	for i := len(values) / 2; i >= 0; i-- {
		h.MinHeapify(i)
	}
	return h
}

func (h *MaxHeap) Insert(key int) {
	h.values = append(h.values, key)
}

func (h *MinHeap) Insert(key int) {
	h.values = append(h.values, key)
}

func (h *MaxHeap) Greater(i, j int) bool {
	return h.values[i] > h.values[j]
}

func (h *MinHeap) Less(i, j int) bool {
	return h.values[i] < h.values[j]
}

func (h *MaxHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *MinHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *MaxHeap) Pop() int {
	old := h.values
	length := len(h.values)
	last := h.values[length-1]
	h.values = old[:length-1]
	return last
}

func (h *MinHeap) Pop() int {
	old := h.values
	length := len(h.values)
	last := h.values[length-1]
	h.values = old[:length-1]
	return last
}

func (h *MaxHeap) MaxHeapify(index int) {
	l, r := 2*index+1, 2*index+2
	max := index

	if l < len(h.values) && h.values[l] > h.values[max] {
		max = l
	}

	if r < len(h.values) && h.values[r] > h.values[max] {
		max = r
	}

	if max != index {
		h.values[index], h.values[max] = h.values[max], h.values[index]
		h.MaxHeapify(max)
	}
}

func (h *MinHeap) MinHeapify(index int) {
	l, r := 2*index+1, 2*index+2
	min := index

	fmt.Println(h.values)
	if l < len(h.values) && h.values[l] < h.values[min] {
		min = l
	}

	if r < len(h.values) && h.values[r] < h.values[min] {
		min = r
	}

	if min != index {
		h.values[index], h.values[min] = h.values[min], h.values[index]
		h.MinHeapify(min)
	}
}

func (h *MaxHeap) Length() int {
	return len(h.values)
}

func (h *MinHeap) Length() int {
	return len(h.values)
}
