package gqueue_test

import (
	"fmt"
	"testing"
	"xiezy1997/gopkg/gqueue"
)

func TestQueue(t *testing.T) {
	arr := []int{1, 5, 3, 4, 2}
	q := gqueue.NewQueue[int]()
	q.Clear()
	for _, v := range arr {
		q.TPush(v)
	}
	i := 0
	for !q.IsEmpty() {
		if i >= 3 {
			q.Clear()
			fmt.Println("queue clear")
		}
		fmt.Println(q.TPop())
		i++
	}
}
