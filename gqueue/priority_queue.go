package gqueue

import "container/heap"

// The `PriQueItem` type is a generic interface in Go that defines a method `Less` for comparing items.
// @property {bool} Less - The `Less` method in the `PriQueItem` interface is used to compare the
// current item with another item of the same type `t`. It should return `true` if the current item is
// considered less than the item being compared to, and `false` otherwise. This method is typically
type PriQueItem[t any] interface {
	Less(b t) bool
}

// The above type defines a priority queue data structure in Go.
// @property {[]t} items - The `items` property in the `PriQueue` struct is a slice that holds elements
// of type `t`, where `t` is a generic type parameter representing the type of elements stored in the
// priority queue.
type PriQueue[t PriQueItem[t]] struct {
	items []t
	len   int
}

// The `NewPriQueue` function is a constructor function for creating a new instance of a priority queue
// (`PriQueue`) in Go. Here's a breakdown of what it does:
func NewPriQueue[t PriQueItem[t]]() IQueue[t] {
	p := &PriQueue[t]{make([]t, 0), 0}
	heap.Init(p)
	return p
}

// func The `func (p PriQueue[t]) Len() int` method is implementing the `Len` function for the `PriQueue`
// type in Go. This method returns the number of elements currently stored in the priority queue `p` by
// returning the length of the `items` slice within the `PriQueue` struct.
func (p PriQueue[t]) Len() int { return p.len }

// The `func (p PriQueue[t]) Swap(i, j int)` method is implementing the `Swap` function for the
// `PriQueue` type in Go. This method is responsible for swapping the elements at positions `i` and `j`
// within the `items` slice of the priority queue `p`.
func (p PriQueue[t]) Swap(i, j int) { p.items[i], p.items[j] = p.items[j], p.items[i] }

// The `func (p PriQueue[t]) Less(i, j int) bool` method is implementing the `Less` function for the
// `PriQueue` type in Go. This method is used to compare the elements at positions `i` and `j` within
// the `items` slice of the priority queue `p`. It calls the `Less` method defined in the `PriQueItem`
// interface for the elements at positions `i` and `j` to determine their relative order in the
// priority queue. The method returns `true` if the element at position `i` is considered less than the
// element at position `j`, and `false` otherwise.
func (p PriQueue[t]) Less(i, j int) bool { return p.items[i].Less(p.items[j]) }

// The `func (p *PriQueue[t]) Push(item any)` method is adding an item to the priority queue `p`.
func (p *PriQueue[t]) Push(item any) {
	p.items = append(p.items, item.(t))
	p.len += 1
}

// The `func (p *PriQueue[t]) Pop() any` method in the provided Go code snippet is implementing the
// functionality to remove and return the top element (with the highest priority) from the priority
// queue `p`.
func (p *PriQueue[t]) Pop() any {
	d := p.items[p.Size()-1]
	p.items = p.items[:p.Size()-1]
	p.len -= 1
	return d
}

func (p *PriQueue[t]) IsEmpty() bool { return p.Len() == 0 }

func (p *PriQueue[t]) Size() int { return p.Len() }

func (p *PriQueue[t]) Clear() {
	p.len = 0
	p.items = p.items[:0]
}

func (p *PriQueue[t]) TPush(item t) { heap.Push(p, item) }

func (p *PriQueue[t]) TPop() (t, bool) {
	if p.IsEmpty() {
		var tmp t
		return tmp, false
	}
	return heap.Pop(p).(t), true
}
