package linklist

import "errors"

var (
	EmptyError = errors.New("queue is empty") // 链表为空
)

type node struct {
	data interface{}
	next *node
}

func newNode() *node {
	return &node{}
}

func (n *node) setData(data interface{}) *node {
	n.data = data
	return n
}

func (n *node) setNext(next *node) *node {
	n.next = next
	return n
}
