package dataStruct

import (
	"fmt"
)

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// 先序
func PreOrder(root *TreeNode) {

	if root == nil {
		return
	}
	fmt.Println(root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// 中序
func MidOrder(root *TreeNode) {

	if root == nil {
		return
	}
	PreOrder(root.Left)
	fmt.Println(root.Data)
	PreOrder(root.Right)
}

// 中序
func PostOrder(root *TreeNode) {

	if root == nil {
		return
	}
	PreOrder(root.Left)
	PreOrder(root.Right)
	fmt.Println(root.Data)
}

// 层序 使用队列
func LayerOrder(root *TreeNode) []string {
	var result []string

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {

		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			result = append(result, node.Data)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]

	}

	return result
}
