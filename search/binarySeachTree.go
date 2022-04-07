package search

import "fmt"

// 二叉查找树
type BinarySearchTree struct {
	Root *BinarySearchTreeNode // 树根节点
}

type BinarySearchTreeNode struct {
	Value int64                 //值
	Timse int64                 //值出现的次数
	Left  *BinarySearchTreeNode //左孩子
	Right *BinarySearchTreeNode //右孩子
}

// 初始化二叉查找树
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

// 添加元素
func (tree *BinarySearchTree) Add(val int64) {
	// 如果没有树根，直接添加根节点
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: val}
		return
	}
	tree.Root.Add(val)
}

func (node *BinarySearchTreeNode) Add(val int64) {
	if val < node.Value {
		// 如果插入的值比节点的值小，那么要插入到该节点的左子树中
		// 如果左子树为空，直接添加
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: val}
		} else {
			// 否则递归
			node.Left.Add(val)
		}
	} else if val < node.Value {
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: val}
		} else {
			node.Right.Add(val)
		}
	} else { //相同就让times++
		node.Timse = node.Timse + 1

	}
}

// 查找最大值或者最小值
// 最小值
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	//如果节点的左子树已经为空说明我们已经找到最左边的节点
	if node.Left == nil {
		return node
	}
	return node.Left.FindMinValue()
}

// 最大值
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	//如果节点的右子树已经为空说明我们已经找到最右边的节点
	if node.Right == nil {
		return node
	}
	return node.Right.FindMinValue()
}

// 查找指定的元素
func (tree *BinarySearchTree) Find(val int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(val)
}

func (node *BinarySearchTreeNode) Find(val int64) *BinarySearchTreeNode {
	if node.Value == val { // 如果该节点刚刚等于该值，那么返回该节点
		return node
	} else if val < node.Value { // 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil { // 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Left.Find(val)

	} else {
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(val)

	}
}

// 查找指定节点父元素
func (tree *BinarySearchTree) FindParent(val int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	if tree.Root.Value == val {
		return nil
	}
	return tree.Root.FindParent(val)
}
func (node *BinarySearchTreeNode) FindParent(val int64) *BinarySearchTreeNode {
	// 外层没有值相等的判定，因为在内层已经判定完毕后返回父亲节点。
	if val < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		leftTree := node.Left
		if leftTree == nil {
			return nil
		}
		// 左子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
		if leftTree.Value == val {
			return node
		} else {
			return leftTree.FindParent(val)
		}
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		rightTree := node.Right
		if rightTree == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		// 右子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
		if rightTree.Value == val {
			return node
		} else {
			return rightTree.FindParent(val)
		}
	}
}

// 删除元素
/*
1. 第一种情况，删除的是根节点，且根节点没有儿子，直接删除即可。
2. 第二种情况，删除的节点有父亲节点，但没有子树，也就是删除的是叶子节点，直接删除
即可。
3. 第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中
的最小元素来替换删除的节点，这时二叉查找树的性质又满足了。右子树的最小元素，只
要一直往右子树的左边一直找一直找就可以找到。
4. 第四种情况，删除的节点只有一个子树，那么该子树直接替换被删除的节点即可。*/

// 删除指定的元素
func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}
	// 查找该值是否存在
	node := tree.Root.Find(value)
	if node == nil {
		// 不存在该值，直接返回
		return
	}
	// 查找该值的父亲节点
	parent := tree.Root.FindParent(value)
	// 第一种情况，删除的是根节点，且根节点没有儿子
	if parent == nil && node.Left == nil && node.Right == nil {
		// 置空后直接返回
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		// 第二种情况，删除的节点有父亲节点，但没有子树
		// 如果删除的是节点是父亲的左儿子，直接将该值删除即可
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			// 删除的原来是父亲的右儿子，直接将该值删除即可
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		// 第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，
		// 那么用右子树中的最小元素来替换删除的节点。
		// 右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到，替换
		// 后二叉查找树的性质又满足了。
		// 找右子树中最小的值，一直往右子树的左边找
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// 把最小的节点删掉
		tree.Delete(minNode.Value)
		// 最小值的节点替换被删除节点
		node.Value = minNode.Value
		node.Timse = minNode.Timse
	} else {
		// 第四种情况，只有一个子树，那么该子树直接替换被删除的节点即可
		// 父亲为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		// 左子树不为空
		if node.Left != nil {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的左子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}

// 中序遍历
func (tree *BinarySearchTree) MidOrder() {
	tree.Root.MidOrder()
}
func (node *BinarySearchTreeNode) MidOrder() {
	if node == nil {
		return
	}
	// 先打印左子树
	node.Left.MidOrder()
	// 按照次数打印根节点
	for i := 0; i <= int(node.Timse); i++ {
		fmt.Println(node.Value)
	}
	// 打印右子树
	node.Right.MidOrder()
}
