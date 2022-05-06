package queue

import (
	"context"
	"sync"
)

type syncQueue struct {
	queue *queue
	mutex *sync.Mutex
}

func (s *syncQueue) Push(ctx context.Context, data interface{}) error {
	s.mutex.Lock()
	err := s.queue.Push(ctx, data)
	s.mutex.Unlock()
	return err
}

func (s *syncQueue) Pop(ctx context.Context) {
	s.mutex.Lock()
	s.queue.Pop(ctx)
	s.mutex.Unlock()
}

func (s *syncQueue) Size(ctx context.Context) int {
	s.mutex.Lock()
	size := s.queue.Size(ctx)
	s.mutex.Unlock()
	return size
}

func (s *syncQueue) Capacity(ctx context.Context) int {
	s.mutex.Lock()
	capacity := s.queue.Capacity(ctx)
	s.mutex.Unlock()
	return capacity
}

func (s *syncQueue) IsEmpty(ctx context.Context) bool {
	s.mutex.Lock()
	empty := s.queue.IsEmpty(ctx)
	s.mutex.Unlock()
	return empty
}

func (s *syncQueue) IsFull(ctx context.Context) bool {
	s.mutex.Lock()
	full := s.queue.IsFull(ctx)
	s.mutex.Unlock()
	return full
}

func (s *syncQueue) Front(ctx context.Context) (interface{}, error) {
	s.mutex.Lock()
	front, err := s.queue.Front(ctx)
	s.mutex.Unlock()
	return front, err
}

func (s *syncQueue) Back(ctx context.Context) (interface{}, error) {
	s.mutex.Lock()
	back, err := s.queue.Back(ctx)
	s.mutex.Unlock()
	return back, err
}

func (s *syncQueue) DataList(ctx context.Context) []interface{} {
	s.mutex.Lock()
	data := s.queue.DataList(ctx)
	s.mutex.Unlock()
	return data
}

func NewSyncQueue(length int) IQueue {
	return &syncQueue{
		queue: &queue{
			startIdx: 0,
			endIdx:   0,
			length:   length,
			dataList: make([]interface{}, length+spareCnt),
		},
		mutex: &sync.Mutex{},
	}
}
