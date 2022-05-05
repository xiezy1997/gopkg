package queue

import (
	"context"
	"fmt"
	"testing"
)

func TestSyncQueue(t *testing.T) {
	q := NewSyncQueue[int](5)
	ctx := context.Background()
	for i := 1; i <= 6; i++ {
		fmt.Printf("queue size=%v ", q.Size(ctx))
		err := q.Push(ctx, i)
		fmt.Println(err)
	}
	fmt.Println(q.Back(ctx))
	for i := 1; i <= 6; i++ {
		fmt.Println(q.Front(ctx))
		q.Pop(ctx)
	}
	for i := 3; i > 0; i-- {
		q.Push(ctx, i)
	}
	fmt.Println(q.Size(ctx))
	fmt.Println(q.DataList(ctx))
	for i := 3; i >= 0; i-- {
		q.Pop(ctx)
		q.Push(ctx, i)
		fmt.Println(q.Size(ctx))
		fmt.Println(q.DataList(ctx))
	}
}
