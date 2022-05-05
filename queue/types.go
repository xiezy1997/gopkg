package queue

import "errors"

const spareCnt = 5 // 空余的队列数量，至少为1

var (
	CapacityError = errors.New("queue capacity error") // 容量不足
	EmptyError    = errors.New("queue is empty")       // 队列为空
)
