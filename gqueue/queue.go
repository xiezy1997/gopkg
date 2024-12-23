package gqueue

type queue[t any] struct {
	items []t
	len   int
}

// NewQueue The NewQueue function creates a new queue with an empty list of items and a length of 0.
//
//	@return *Queue
func NewQueue[t any]() IQueue[t] {
	return &queue[t]{make([]t, 0), 0}
}

func (q *queue[t]) IsEmpty() bool { return q.len == 0 }

func (q *queue[t]) Size() int { return q.len }

func (q *queue[t]) Clear() {
	q.items = q.items[:0]
	q.len = 0
}

func (q *queue[t]) TPush(item t) {
	q.items = append(q.items, item)
	q.len += 1
}

func (q *queue[t]) TPop() (t, bool) {
	if q.IsEmpty() {
		var tmp t
		return tmp, false
	}
	old := q.items[0]
	if q.Size() > 1 {
		q.items = q.items[1:]
	} else {
		q.items = make([]t, 0)
	}
	q.len -= 1
	return old, true
}
