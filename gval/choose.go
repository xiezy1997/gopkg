package gval

// The `If` function in Go returns either the true value or the false value based on a given condition.
func If[t any](cond bool, trueVal, falseVal t) t {
	if cond {
		return trueVal
	}
	return falseVal
}

// The `IfLazy` function in Go takes a condition and two functions as arguments, and lazily evaluates
// and returns the result of one of the functions based on the condition.
func IfLazy[t any](cond bool, trueFunc, falseFunc func() t) t {
	if cond {
		return trueFunc()
	}
	return falseFunc()
}

// The function `IfLazyT` returns the result of `trueFunc` if `cond` is true, otherwise it returns
// `falseVal`.
func IfLazyT[t any](cond bool, trueFunc func() t, falseVal t) t {
	if cond {
		return trueFunc()
	}
	return falseVal
}

// The function `IfLazyF` returns either a value or the result of a function based on a condition.
func IfLazyF[t any](cond bool, trueVal t, falseFunc func() t) t {
	if cond {
		return trueVal
	}
	return falseFunc()
}
