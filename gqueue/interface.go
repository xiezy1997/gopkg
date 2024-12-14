package gqueue

type IQueue[t any] interface {

	// IsEmpty func This function `func IsEmpty() bool` is a method defined on the `Queue` struct in Go.
	// It checks if the queue is empty by comparing the length of the queue (`q.len`) to 0. If the length
	// is 0, it returns `true`, indicating that the queue is empty; otherwise, it returns `false`,
	// indicating that the queue is not empty.
	//
	// @return bool
	IsEmpty() bool

	// Size func The `func Size() int` method in the provided Go code is a method defined on the
	// `Queue` struct. This method is used to retrieve the current size of the queue, which is represented
	// by the length of the queue (`q.len`).
	//
	// @return int
	Size() int

	// Clear The `Clear` method in the `Queue` struct is used to reset the queue to an empty state. It achieves
	// this by setting the `items` slice to a new slice with a length of 0, effectively removing all
	// elements from the queue. Additionally, it sets the `len` field of the queue to 0, indicating that
	// the queue is now empty. This method is useful when you want to clear the contents of the queue
	// without creating a new queue instance.
	Clear()

	// TPush func The `func TPush(item t)` method in the provided Go code is responsible for adding an
	// item to the queue.
	//
	// @param item
	TPush(item t)

	// TPop func The `func TPop() (t, bool)` method in the provided Go code is responsible for popping
	// an item from the queue.
	//
	// @return item
	// @return success
	TPop() (t, bool)
}
