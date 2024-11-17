package gulc

import (
	"fmt"
)


type LinkedListNode[T any] struct  {
	Val T
	Next *LinkedListNode[T]
}

type LinkedList[T any] struct {
	Head *LinkedListNode[T]
}


// ReverseEveryKNodeV1 每隔K个链表翻转，最后不满的一部分也翻转
func(l *LinkedList[T]) ReverseEveryKNodeV1(k int) {
	head := &LinkedListNode[T]{}
    before := head
	tr := l.Head
	for tr != nil {
		count := 0
        cur := tr
        for tr.Next != nil && count < k - 1 {
            count++
            tr = tr.Next
        }
        temp := tr.Next
        tr.Next = nil
        tr = temp
        list := &LinkedList[T]{Head: cur}
        // 基本思路就是先切割，然后一段一段翻转后连接起来
        list.Reverse()
        before.Next = list.Head
        before = cur
	}
    l.Head = head.Next
}

// ReverseEveryKNodeV2 每隔K个链表翻转，最后不满的一部分不翻转
func(l *LinkedList[T]) ReverseEveryKNodeV2(k int) {
	head := &LinkedListNode[T]{}
    before := head
	tr := l.Head
	for tr != nil {
		count := 0
        cur := tr
        for tr.Next != nil && count < k - 1 {
            count++
            tr = tr.Next
        }
        temp := tr.Next
        tr.Next = nil
        tr = temp
        list := &LinkedList[T]{Head: cur}
        // 基本思路就是先切割，然后一段一段翻转后连接起来
        if count == k - 1 {   // 只有一整片才翻转
            list.Reverse()
        }
        before.Next = list.Head
        before = cur
	}
    l.Head = head.Next
}


// Reverse 翻转整个链表
func (l *LinkedList[T]) Reverse() {
    tr := l.Head.Next
    after := l.Head
    l.Head.Next = nil
    for tr != nil {
        temp := tr.Next
        tr.Next = after
        after = tr
        tr = temp
    }
    l.Head = after
}


// Print 打印链表
func (l *LinkedList[T]) ToString() string {
	tr := l.Head
	ans := ""
	if tr != nil {
		ans += fmt.Sprintf("%+v", tr.Val)
	}
	tr = tr.Next
	for tr != nil {
		ans += fmt.Sprintf(" <-> %+v", tr.Val)
		tr = tr.Next
	}
	return ans
}


// NewLinkedList 根据数组新建一个链表
func NewLinkedList[T any] (arr []T) *LinkedList[T] {
	head := LinkedListNode[T]{Val: arr[0], Next: nil}
	tr := &head
	for i := 1; i < len(arr); i++ {
		tr.Next = &LinkedListNode[T]{Val: arr[i], Next: nil}
		tr = tr.Next
	}
	return &LinkedList[T]{Head: &head}
}


func (l *LinkedList[T]) InsertNode(v T) {
	curNode := &LinkedListNode[T]{Val: v}
	curNode.Next = l.Head
	l.Head = curNode
}


// ListNode 表示一个双向链表中的一个节点
type ListNode[T any] struct {
	Key  int
	Val  T
    Freq int
	Next *ListNode[T]
	Prev *ListNode[T]
}

func NewListNode[T any](key int, val T) *ListNode[T] {
	return &ListNode[T]{Key: key, Val: val}
}

// List 定义一个双向链表
type List[T any] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
}

// NewList 创建一个空的双向链表，并且使用伪头部和伪尾部
func NewList[T any]() *List[T] {
	head := &ListNode[T]{Key: -1}
	tail := &ListNode[T]{Key: -1}
	head.Next = tail
	tail.Prev = head
	return &List[T]{Head: head, Tail: tail}
}

// AppendNodeHead 将某个节点插在链表头部
func (l *List[T]) AppendNodeHead(node *ListNode[T]) {
	node.Next = l.Head.Next
	l.Head.Next.Prev = node
	node.Prev = l.Head
	l.Head.Next = node
}

// AppendNodeTail 将某个节点插在链表尾部
func (l *List[T]) AppendNodeTail(node *ListNode[T]) {
	l.Tail.Prev.Next = node
	node.Prev = l.Tail.Prev
	node.Next = l.Tail
	l.Tail.Prev = node
}

// DeleteNode 删除一个指定的节点
func (l *List[T]) DeleteNode(node *ListNode[T]) {
	// 避免panic
	if node.Prev == nil || node.Next == nil {
		fmt.Println("error: illegal node to be deleted")
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// IsEmpty 判断一个链表是否为空
func (l *List[T]) IsEmpty() bool {
    return l.Head.Next == l.Tail
}

func (l *List[T]) RemoveHead() T {
	node := l.Head.Next;
	l.DeleteNode(node);
	return node.Val;
}

func (l *List[T]) AppendTail(v T) {
	node := NewListNode(-1, v)
	l.AppendNodeHead(node)
}

func (l *List[T]) Begin() *ListNode[T] {
	return l.Head.Next
}

func (l *List[T]) End() *ListNode[T] {
	return l.Tail
}
