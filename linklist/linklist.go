package linklist

import (
	"context"
)

type LinkList struct {
	head *node
	tail *node
	cnt  int
}

func (l *LinkList) Head(ctx context.Context) (interface{}, error) {
	if l.cnt == 0 {
		return nil, EmptyError
	}
	return l.head.data, nil
}

func (l *LinkList) Tail(ctx context.Context) (interface{}, error) {
	if l.cnt == 0 {
		return nil, EmptyError
	}
	return l.tail.data, nil
}

func (l *LinkList) InsertToHead(ctx context.Context, data interface{}) {
	node := newNode().setData(data).setNext(l.head)
	l.head = node
	l.cnt++
	if l.tail == nil {
		l.tail = node
		return
	}
}

func (l *LinkList) InsertToTail(ctx context.Context, data interface{}) {
	node := newNode().setData(data)
	l.cnt++
	if l.cnt == 1 {
		l.head = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *LinkList) DeleteFromHead(ctx context.Context) error {
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

func (l *LinkList) DeleteNode(ctx context.Context, delNode interface{}, eqFunc func(dataNode, delNode interface{}) bool) error {
	var p, pre, head, tail *node
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

func (l *LinkList) Size(ctx context.Context) int {
	return l.cnt
}

func (l *LinkList) IsEmpty(ctx context.Context) bool {
	return l.cnt == 0
}

func (l *LinkList) Clear(ctx context.Context) {
	l.init()
}

func (l *LinkList) Iterator(ctx context.Context) func() (interface{}, bool) {
	p := l.head
	return func() (interface{}, bool) {
		if p == nil {
			return nil, true
		}
		data := p.data
		p = p.next
		return data, false
	}
}

func (l *LinkList) init() {
	l.tail = nil
	l.head = nil
	l.cnt = 0
}

func NewLinkList() ILinkList {
	return &LinkList{}
}
