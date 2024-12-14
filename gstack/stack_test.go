package gstack_test

import (
	"fmt"
	"testing"
	"xiezy1997/gopkg/gstack"
)


func TestStack(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	s := gstack.NewStack[int]()
	for _, v := range arr {
		s.TPush(v)
	}

	for !s.IsEmpty() {
		fmt.Println(s.TPop())
	}
}