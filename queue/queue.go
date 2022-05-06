package queue

import "context"

type queue struct {
	startIdx int
	endIdx   int
	length   int
	dataList []interface{}
}

func (q *queue) Push(ctx context.Context, data interface{}) error {
	if q.IsFull(ctx) {
		return CapacityError
	}
	q.dataList[q.endIdx] = data
	q.endIdx = (q.endIdx + 1) % (q.length + spareCnt)
	return nil
}

func (q *queue) Pop(ctx context.Context) {
	if q.IsEmpty(ctx) {
		return
	}
	q.startIdx = (q.startIdx + 1) % (q.length + spareCnt)
}

func (q *queue) Size(ctx context.Context) int {
	if q.startIdx <= q.endIdx {
		return q.endIdx - q.startIdx
	}
	return q.length + spareCnt - q.startIdx + q.endIdx
}

func (q *queue) Capacity(ctx context.Context) int {
	return q.length
}

func (q *queue) IsEmpty(ctx context.Context) bool {
	return q.startIdx == q.endIdx
}

func (q *queue) IsFull(ctx context.Context) bool {
	return q.Size(ctx) == q.length
}

func (q *queue) Front(ctx context.Context) (interface{}, error) {
	if q.IsEmpty(ctx) {
		return nil, EmptyError
	}
	return q.dataList[q.startIdx], nil
}

func (q *queue) Back(ctx context.Context) (interface{}, error) {
	if q.IsEmpty(ctx) {
		return nil, EmptyError
	}
	return q.dataList[q.endIdx-1], nil
}

func (q *queue) DataList(ctx context.Context) []interface{} {
	if q.IsEmpty(ctx) {
		return []interface{}{}
	}
	if q.startIdx < q.endIdx {
		return q.dataList[q.startIdx:q.endIdx]
	}
	return append(q.dataList[q.startIdx:q.length+spareCnt], q.dataList[:q.endIdx]...)
}

func NewQueue(length int) IQueue {
	return &queue{
		startIdx: 0,
		endIdx:   0,
		length:   length,
		dataList: make([]interface{}, length+spareCnt),
	}
}
