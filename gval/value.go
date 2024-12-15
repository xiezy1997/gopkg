package gval

// The `Of` function in Go returns a pointer to the input value of any type.
func Of[t any](a t) *t {
	return &a
}

// The function `Val` returns the value pointed to by a pointer, or an empty value if the pointer is
// nil.
func Val[t any](a *t) t {
	if a == nil {
		return Empty[t]()
	}
	return *a
}

// The function `Empty` in Go returns an empty value of any type provided as a generic parameter.
func Empty[t any]() t {
	var tmp t
	return tmp
}