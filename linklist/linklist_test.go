package linklist

import (
	"context"
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	list := NewLinkList[int]()
	ctx := context.Background()
	fmt.Println("初始化")
	fmt.Println(list.IsEmpty(ctx))
	fmt.Println(list.Size(ctx))
	iter := list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}

	fmt.Printf("\n塞入2134\n")
	list.InsertToTail(ctx, 1)
	list.InsertToHead(ctx, 2)
	list.InsertToTail(ctx, 3)
	list.InsertToTail(ctx, 4)
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}

	fmt.Printf("\n删除头节点2\n")
	_ = list.DeleteFromHead(ctx)
	fmt.Println(list.IsEmpty(ctx))
	fmt.Println(list.Size(ctx))
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}

	fmt.Printf("\n删除头节点1\n")
	_ = list.DeleteFromHead(ctx)
	fmt.Println(list.IsEmpty(ctx))
	fmt.Println(list.Size(ctx))
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}

	fmt.Printf("\n删除头节点3\n")
	_ = list.DeleteFromHead(ctx)
	fmt.Println(list.IsEmpty(ctx))
	fmt.Println(list.Size(ctx))
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}

	fmt.Printf("\n删除头节点4\n")
	_ = list.DeleteFromHead(ctx)
	fmt.Println(list.IsEmpty(ctx))
	fmt.Println(list.Size(ctx))
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}

	fmt.Printf("\n删空链表\n")
	fmt.Println(list.IsEmpty(ctx))
	fmt.Println(list.Size(ctx))

	fmt.Printf("\n塞入21342\n")
	list.InsertToTail(ctx, 1)
	list.InsertToHead(ctx, 2)
	list.InsertToTail(ctx, 3)
	list.InsertToTail(ctx, 4)
	list.InsertToTail(ctx, 2)
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}
	fmt.Printf("\n删除节点2\n")
	_ = list.DeleteNode(ctx, 2, func(delNode, dataNode int) bool {
		if delNode == dataNode {
			return true
		}
		return false
	})
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}
	fmt.Printf("\n删除节点3\n")
	_ = list.DeleteNode(ctx, 3, func(delNode, dataNode int) bool {
		if delNode == dataNode {
			return true
		}
		return false
	})
	iter = list.Iterator(ctx)
	for {
		data, end := iter()
		if end {
			break
		}
		fmt.Printf("%v ", data)
	}
	fmt.Printf("\n")
	list.Clear(ctx)
	fmt.Println(list)
}
