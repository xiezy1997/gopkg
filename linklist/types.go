package linklist

import "errors"

var (
	EmptyError = errors.New("queue is empty") // 链表为空
)

type node[T any] struct {
	data T
	next *node[T]
}

func newNode[T any]() *node[T] {
	return &node[T]{}
}

func (n *node[T]) setData(data T) *node[T] {
	n.data = data
	return n
}

func (n *node[T]) setNext(next *node[T]) *node[T] {
	n.next = next
	return n
}
