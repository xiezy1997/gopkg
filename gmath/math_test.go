package gmath_test

import (
	"fmt"
	"testing"

	"xiezy1997/gopkg/gmath"
)

func Test(t *testing.T) {
	fmt.Println(gmath.Max(1, 2))
	fmt.Println(gmath.Max(1.2, 1.3))
	fmt.Println(gmath.Max("a", "ab", "ac"))

	fmt.Println(gmath.Min(1, 2))
	fmt.Println(gmath.Min(1.2, 1.3))
	fmt.Println(gmath.Min("a", "ab", "ac"))

	fmt.Println(gmath.IdxMax(1, 2))
	fmt.Println(gmath.IdxMax(1.2, 1.3))
	fmt.Println(gmath.IdxMax("a", "ab", "ac"))

	fmt.Println(gmath.IdxMin(1, 2))
	fmt.Println(gmath.IdxMin(1.2, 1.3))
	fmt.Println(gmath.IdxMin("a", "ab", "ac"))

	fmt.Println(gmath.Abs(1))
	fmt.Println(gmath.Abs(-1))

	fmt.Println(gmath.Mod(5, 3))
	fmt.Println(gmath.Mod(-5, 3))

	fmt.Println(gmath.Pwd(3, 3))
	fmt.Println(gmath.PwdMod(3, 3, 5))

}
