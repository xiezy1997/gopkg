package queue

import (
	"context"
	"sync"
)

type syncQueue[T any] struct {
	queue *queue[T]
	mutex *sync.Mutex
}

func (s *syncQueue[T]) Push(ctx context.Context, data T) error {
	s.mutex.Lock()
	err := s.queue.Push(ctx, data)
	s.mutex.Unlock()
	return err
}

func (s *syncQueue[T]) Pop(ctx context.Context) {
	s.mutex.Lock()
	s.queue.Pop(ctx)
	s.mutex.Unlock()
}

func (s *syncQueue[T]) Size(ctx context.Context) int {
	s.mutex.Lock()
	size := s.queue.Size(ctx)
	s.mutex.Unlock()
	return size
}

func (s *syncQueue[T]) Capacity(ctx context.Context) int {
	s.mutex.Lock()
	capacity := s.queue.Capacity(ctx)
	s.mutex.Unlock()
	return capacity
}

func (s *syncQueue[T]) IsEmpty(ctx context.Context) bool {
	s.mutex.Lock()
	empty := s.queue.IsEmpty(ctx)
	s.mutex.Unlock()
	return empty
}

func (s *syncQueue[T]) IsFull(ctx context.Context) bool {
	s.mutex.Lock()
	full := s.queue.IsFull(ctx)
	s.mutex.Unlock()
	return full
}

func (s *syncQueue[T]) Front(ctx context.Context) (T, error) {
	s.mutex.Lock()
	front, err := s.queue.Front(ctx)
	s.mutex.Unlock()
	return front, err
}

func (s *syncQueue[T]) Back(ctx context.Context) (T, error) {
	s.mutex.Lock()
	back, err := s.queue.Back(ctx)
	s.mutex.Unlock()
	return back, err
}

func (s *syncQueue[T]) DataList(ctx context.Context) []T {
	s.mutex.Lock()
	data := s.queue.DataList(ctx)
	s.mutex.Unlock()
	return data
}

func NewSyncQueue[T any](length int) IQueue[T] {
	return &syncQueue[T]{
		queue: &queue[T]{
			startIdx: 0,
			endIdx:   0,
			length:   length,
			dataList: make([]T, length+spareCnt),
		},
		mutex: &sync.Mutex{},
	}
}
