package dataStruct

import "sync"

type LinkStack struct {
	root *LinkNode  // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 头插法入栈，新入栈的节点在头部
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 如果栈顶为空，那么增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否则新元素插入链表的头部
		// 原来的链表
		preNode := stack.root
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v
		// 原来的链表链接到新元素后面
		newNode.Next = preNode
		// 将新节点放在头部
		stack.root = newNode
	}
	// 栈中元素数量+1
	stack.size = stack.size + 1
}

// 出栈
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}
	// 顶部元素要出栈
	topNode := stack.root
	v := topNode.Value
	// 将顶部元素的后继链接链上
	stack.root = topNode.Next
	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v
}

func (stack *LinkStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}

	return stack.root.Value
}

// 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
