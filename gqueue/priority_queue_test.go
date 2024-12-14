package gqueue_test

import (
	"fmt"
	"testing"
	"xiezy1997/gopkg/gqueue"
)

type Item struct {
	idx, val int
}

func (a *Item) Less(b *Item) bool {
	if b.val == a.val {
		return a.idx < b.idx
	}
	return a.val < b.val
}

func TestPriQueue(t *testing.T) {
	q := gqueue.NewPriQueue[*Item]()
	q.Clear()
	arr := []int{3, 2, 5, 1}
	for i, v := range arr {
		q.TPush(&Item{idx: i, val: v})
	}
	i := 0
	for !q.IsEmpty() {
		if i >= 1 {
			q.Clear()
			fmt.Println("queue clear")
		}
		t, b := q.TPop()
		if b {
			fmt.Println(b, t.idx, t.val)
		} else {
			fmt.Println(b)
		}
		i++
	}
}