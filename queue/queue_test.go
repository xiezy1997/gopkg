package queue

import (
	"context"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int](5)
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

func TestStructQueue(t *testing.T) {
	type Task struct {
		id int
	}
	q := NewQueue[*Task](5)
	ctx := context.Background()
	for i := 1; i <= 6; i++ {
		fmt.Printf("queue size=%v ", q.Size(ctx))
		err := q.Push(ctx, &Task{id: i})
		fmt.Println(err)
	}
	fmt.Println(q.Back(ctx))
	for i := 1; i <= 6; i++ {
		fmt.Println(q.Front(ctx))
		q.Pop(ctx)
	}
	for i := 3; i > 0; i-- {
		q.Push(ctx, &Task{id: i})
	}
	fmt.Println(q.Size(ctx))
	fmt.Println(q.DataList(ctx))
	for i := 3; i >= 0; i-- {
		q.Pop(ctx)
		q.Push(ctx, &Task{id: i})
		fmt.Println(q.Size(ctx))
		fmt.Println(q.DataList(ctx))
	}
}
