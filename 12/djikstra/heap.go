package djikstra

import (
	"sync"
)

type Heap struct {
	elements []*Point
	mutex    sync.RWMutex
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

// push an element to the heap, re-arrange the heap
func (h *Heap) Push(element *Point) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
}

// pop the top of the heap, which is the min DistanceFromStart
func (h *Heap) Pop() (i *Point) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	index := h.findBest()
	i = h.elements[index]
	h.elements[index] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	return
}

func (h *Heap) findBest() int {
	best := 0
	bestDistance := float64(999)
	for i, p := range h.elements {
		if p.DistanceFromStart < bestDistance {
			best = i
			bestDistance = p.DistanceFromStart
		}
	}
	return best
}

// rearrange the heap
func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].DistanceFromStart < h.elements[smallest].DistanceFromStart {
		smallest = left
	}
	if right < size && h.elements[right].DistanceFromStart < h.elements[smallest].DistanceFromStart {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
