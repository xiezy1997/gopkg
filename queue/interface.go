// Package queue
// Golang version: 1.18.1
// Description: 在golang中，标准库里没有队列，该包是以C++队列的常用方法定义接口，实现通用的队列方法
package queue

import "context"

type IQueue[T any] interface {
	// Push 在队尾插入一个元素
	//  @param ctx
	//  @param data
	//  @return error
	Push(ctx context.Context, data T) error
	// Pop 将队列中最靠前位置的元素删除
	//  @param ctx
	Pop(ctx context.Context)

	// Size 返回队列中元素个数
	//  @param ctx
	//  @return int
	Size(ctx context.Context) int
	// Capacity 返回队列容量
	//  @param ctx
	//  @return int
	Capacity(ctx context.Context) int
	// IsEmpty 如果队列空则返回true
	//  @param ctx
	//  @return bool
	IsEmpty(ctx context.Context) bool
	// IsFull 如果队列满了则返回true
	//  @param ctx
	//  @return bool
	IsFull(ctx context.Context) bool

	// Front 返回队列中的第一个元素
	//  @param ctx
	//  @return T
	//  @return error
	Front(ctx context.Context) (T, error)
	// Back 返回队列中最后一个元素
	//  @param ctx
	//  @return T
	//  @return error
	Back(ctx context.Context) (T, error)
	// DataList 获取数据列表
	//  @param ctx
	//  @return []T
	DataList(ctx context.Context) []T
}
