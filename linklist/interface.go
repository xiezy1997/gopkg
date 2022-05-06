// Package linklist
// Golang version: 1.17
// Description: 使用golang实现链表，支持头尾插入，头尾取出、头删除，以及迭代器函数操作等，出参类型为interface，则需要外部断言
package linklist

import "context"

type ILinkList interface {
	// Head 获取链表第一个元素
	//  @param ctx
	//  @return interface{}: 如果链表为空，则返回nil
	//  @return error
	Head(ctx context.Context) (interface{}, error)
	// Tail 获取链表最后一个元素
	//  @param ctx
	//  @return interface{}: 如果链表为空，则返回nil
	//  @return error
	Tail(ctx context.Context) (interface{}, error)

	// InsertToHead 插入一个元素到头部
	//  @param ctx
	//  @param data
	InsertToHead(ctx context.Context, data interface{})
	// InsertToTail 插入一个元素到尾部
	//  @param ctx
	//  @param data
	InsertToTail(ctx context.Context, data interface{})

	// DeleteFromHead 删除第一个节点
	//  @param ctx
	//  @return error
	DeleteFromHead(ctx context.Context) error
	// DeleteNode 删除等值的节点
	//  @param ctx
	//  @param delNode: 要删除的节点值(T类型)
	//  @param eqFunc: 等值判断函数，返回true为等值，则会删除
	//  @return error
	DeleteNode(ctx context.Context, delNode interface{}, eqFunc func(dataNode, delNode interface{}) bool) error

	// Size 获取链表长度
	//  @param ctx
	//  @return int
	Size(ctx context.Context) int
	// IsEmpty 判断链表是否为空
	//  @param ctx
	//  @return bool
	IsEmpty(ctx context.Context) bool
	// Clear 清空链表
	//  @param ctx
	Clear(ctx context.Context)

	// Iterator 获取迭代器函数
	//  @param ctx
	//  @return func() (interface{}, bool): data, isNull
	Iterator(ctx context.Context) func() (interface{}, bool)
}
