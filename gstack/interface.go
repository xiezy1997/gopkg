package gstack

type IStahk[t any] interface {

	// IsEmpty The `IsEmpty` method in the `Stack` struct is checking if the stack is empty by calling the `Size`
	// method and comparing the size of the stack to 0. If the size is 0, it means the stack is empty, and
	// the method returns `true`. Otherwise, it returns `false`.
	// @return bool
	IsEmpty() bool

	// Size The `Size()` method in the `Stack` struct is a method that returns the current size of the stack. It
	// simply returns the value of the `len` field in the `Stack` struct, which represents the number of
	// elements currently in the stack.
	// @return int
	Size() int

	// Clear The `Clear()` method in the `Stack` struct is a function that clears all elements from the stack. It
	// achieves this by resetting the `items` slice to an empty slice (with length 0) and setting the `len`
	// field of the stack to 0. This effectively removes all elements from the stack, making it empty while
	// maintaining the underlying data structure for potential future use.
	Clear()

	// TPush The `TPush` method in the `Stack` struct is used to push an element onto the stack. It takes an item
	// of type `t` as a parameter and appends it to the `items` slice of the stack. Additionally, it
	// increments the `len` field of the stack by 1 to reflect the addition of the new element. This
	// operation effectively adds the provided item to the top of the stack, expanding the stack size by
	// one element.
	// @param item
	TPush(item t)

	// TPop The `TPop` method in the `Stack` struct is used to pop an element from the top of the stack.
	// @return item
	// @return success
	TPop() (t, bool)
}
