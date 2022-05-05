package linklist

import (
	"context"
)

type LinkList[T any] struct {
	head *node[T]
	tail *node[T]
	cnt  int
}

func (l *LinkList[T]) Head(ctx context.Context) (T, error) {
	if l.cnt == 0 {
		var data T
		return data, EmptyError
	}
	return l.head.data, nil
}

func (l *LinkList[T]) Tail(ctx context.Context) (T, error) {
	if l.cnt == 0 {
		var data T
		return data, EmptyError
	}
	return l.tail.data, nil
}

func (l *LinkList[T]) InsertToHead(ctx context.Context, data T) {
	node := newNode[T]().setData(data).setNext(l.head)
	l.head = node
	l.cnt++
	if l.tail == nil {
		l.tail = node
		return
	}
}

func (l *LinkList[T]) InsertToTail(ctx context.Context, data T) {
	node := newNode[T]().setData(data)
	l.cnt++
	if l.cnt == 1 {
		l.head = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *LinkList[T]) DeleteFromHead(ctx context.Context) error {
	if l.cnt == 0 {
		return EmptyError
	}
	l.cnt--
	l.head = l.head.next
	if l.cnt == 0 {
		l.init()
		return nil
	}
	return nil
}

func (l *LinkList[T]) DeleteNode(ctx context.Context, delNode T, eqFunc func(dataNode, delNode T) bool) error {
	var p, pre, head, tail *node[T]
	for p = l.head; p != nil; p = p.next {
		if eqFunc(p.data, delNode) {
			l.cnt--
			if pre == nil {
				continue
			}
			pre.next = p.next
			continue
		}
		if head == nil {
			head = p
		}
		pre = p
		tail = p
	}
	l.head = head
	l.tail = tail
	return nil
}

func (l *LinkList[T]) Size(ctx context.Context) int {
	return l.cnt
}

func (l *LinkList[T]) IsEmpty(ctx context.Context) bool {
	return l.cnt == 0
}

func (l *LinkList[T]) Clear(ctx context.Context) {
	l.init()
}

func (l *LinkList[T]) Iterator(ctx context.Context) func() (T, bool) {
	p := l.head
	return func() (T, bool) {
		if p == nil {
			var t T
			return t, true
		}
		data := p.data
		p = p.next
		return data, false
	}
}

func (l *LinkList[T]) init() {
	l.tail = nil
	l.head = nil
	l.cnt = 0
}

func NewLinkList[T any]() ILinkList[T] {
	return &LinkList[T]{}
}
