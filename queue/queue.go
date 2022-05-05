package queue

import "context"

type queue[T any] struct {
	startIdx int
	endIdx   int
	length   int
	dataList []T
}

func (q *queue[T]) Push(ctx context.Context, data T) error {
	if q.IsFull(ctx) {
		return CapacityError
	}
	q.dataList[q.endIdx] = data
	q.endIdx = (q.endIdx + 1) % (q.length + spareCnt)
	return nil
}

func (q *queue[T]) Pop(ctx context.Context) {
	if q.IsEmpty(ctx) {
		return
	}
	q.startIdx = (q.startIdx + 1) % (q.length + spareCnt)
}

func (q *queue[T]) Size(ctx context.Context) int {
	if q.startIdx <= q.endIdx {
		return q.endIdx - q.startIdx
	}
	return q.length + spareCnt - q.startIdx + q.endIdx
}

func (q *queue[T]) Capacity(ctx context.Context) int {
	return q.length
}

func (q *queue[T]) IsEmpty(ctx context.Context) bool {
	return q.startIdx == q.endIdx
}

func (q *queue[T]) IsFull(ctx context.Context) bool {
	return q.Size(ctx) == q.length
}

func (q *queue[T]) Front(ctx context.Context) (T, error) {
	if q.IsEmpty(ctx) {
		var t T
		return t, EmptyError
	}
	return q.dataList[q.startIdx], nil
}

func (q *queue[T]) Back(ctx context.Context) (T, error) {
	if q.IsEmpty(ctx) {
		var t T
		return t, EmptyError
	}
	return q.dataList[q.endIdx-1], nil
}

func (q *queue[T]) DataList(ctx context.Context) []T {
	if q.IsEmpty(ctx) {
		return []T{}
	}
	if q.startIdx < q.endIdx {
		return q.dataList[q.startIdx:q.endIdx]
	}
	data := make([]T, q.Size(ctx))
	idx := 0
	for i := q.startIdx; i < q.length+spareCnt; i++ {
		data[idx] = q.dataList[i]
		idx++
	}
	for i := 0; i < q.endIdx; i++ {
		data[idx] = q.dataList[i]
		idx++
	}
	return data
}

func NewQueue[T any](length int) IQueue[T] {
	return &queue[T]{
		startIdx: 0,
		endIdx:   0,
		length:   length,
		dataList: make([]T, length+spareCnt),
	}
}
