package gmath

func Min[t ComparableType](a t, b ...t) t {
	for _, v := range b {
		if a > v {
			a = v
		}
	}
	return a
}

func Max[t ComparableType](a t, b ...t) t {
	for _, v := range b {
		if a < v {
			a = v
		}
	}
	return a
}

func IdxMin[t ComparableType](a t, b ...t) (int, t) {
	idx := 0
	for i, v := range b {
		if a > v {
			a, idx = v, i+1
		}
	}
	return idx, a
}

func IdxMax[t ComparableType](a t, b ...t) (int, t) {
	idx := 0
	for i, v := range b {
		if a < v {
			a, idx = v, i+1
		}
	}
	return idx, a
}

func Abs[t NumberType](a t) t {
	if a > 0 {
		return a
	}
	return -a
}

func Mod[t IntType](num, mod t) t {
	if num < 0 {
		return (num%mod + mod) % mod
	}
	return num % mod
}

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
