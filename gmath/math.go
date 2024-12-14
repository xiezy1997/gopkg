package gmath

type ComparableType interface {
	NumberType | string
}

type NumberType interface {
	IntType | FloatType
}

type FloatType interface {
	float32 | float64
}

type IntType interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

// Min The Min function in Go returns the smallest value among the input values.
// @param v0
// @param vs
// @return min value
func Min[t ComparableType](v0 t, vs ...t) t {
	for _, v := range vs {
		if v0 > v {
			v0 = v
		}
	}
	return v0
}

// Max The Max function in Go returns the maximum value among a variable number of input values.
// @param v0
// @param vs
// @return max value
func Max[t ComparableType](v0 t, vs ...t) t {
	for _, v := range vs {
		if v0 < v {
			v0 = v
		}
	}
	return v0
}

// IdxMin The function IdxMin finds the index and value of the minimum element in a list of comparable
// elements.
// @param v0
// @param vs
// @return min value index
// @return min value
func IdxMin[t ComparableType](v0 t, vs ...t) (int, t) {
	idx := 0
	for i, v := range vs {
		if v0 > v {
			v0, idx = v, i+1
		}
	}
	return idx, v0
}

// IdxMax The function IdxMax finds the index and value of the maximum element in a slice of comparable types.
// @param v0
// @param vs
// @return max value index
// @return max value
func IdxMax[t ComparableType](v0 t, vs ...t) (int, t) {
	idx := 0
	for i, v := range vs {
		if v0 < v {
			v0, idx = v, i+1
		}
	}
	return idx, v0
}

// Abs The function Abs calculates the absolute value of a given number.
// @param num
// @return abs(num)
func Abs[t NumberType](num t) t {
	if num > 0 {
		return num
	}
	return -num
}

// Mod The Mod function calculates the modulo operation for two integers, handling negative numbers
// appropriately.
// @param num
// @param mod
// @return num % mod
func Mod[t IntType](num, mod t) t {
	if num < 0 {
		return (num%mod + mod) % mod
	}
	return num % mod
}

// Pwd The function `Pwd` calculates the power of a given base to a specified exponent using an optimized
// algorithm.
// @param base
// @param exp
// @return base ^ exp
func Pwd[t IntType](base t, exp int) t {
	res := t(1)
	for ; exp > 0; exp /= 2 {
		if exp%2 > 0 {
			res = res * base
		}
		base = base * base
	}
	return res
}

// PwdMod The function `PwdMod` calculates the modular exponentiation of a base raised to an exponent with a
// given modulus.
// @param base
// @param exp
// @param mod
// @return (base ^ exp) % mod
func PwdMod[t IntType](base t, exp int, mod t) t {
	res := t(1)
	for ; exp > 0; exp /= 2 {
		if exp%2 > 0 {
			res = res * base % mod
		}
		base = base * base % mod
	}
	return res
}
