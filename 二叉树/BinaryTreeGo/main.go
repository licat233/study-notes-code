package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Key        int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

type BinaryTreeNode struct {
	Root     *TreeNode
	Keys     []int
	Callback func(key int)
}

//构建节点
func (b *BinaryTreeNode) NewBinaryTreeNode(key int) *TreeNode {
	node := &TreeNode{
		Key:        key,
		LeftChild:  nil,
		RightChild: nil,
	}
	return node
}

//插入节点
func (b *BinaryTreeNode) InsertBinaryTreeNode(node *TreeNode, newNode *TreeNode) {
	if newNode.Key < node.Key {
		if node.LeftChild != nil {
			b.InsertBinaryTreeNode(node.LeftChild, newNode)
		} else {
			node.LeftChild = newNode
		}
	} else {
		if node.RightChild != nil {
			b.InsertBinaryTreeNode(node.RightChild, newNode)
		} else {
			node.RightChild = newNode
		}
	}
}

//构建二叉树
func (b *BinaryTreeNode) BuildBinaryTree() {
	for _, key := range b.Keys {
		node := b.NewBinaryTreeNode(key)
		if b.Root == nil {
			b.Root = node
		} else {
			b.InsertBinaryTreeNode(b.Root, node)
		}
	}
}

//遍历二叉树：前序遍历：中->左->右
func (b *BinaryTreeNode) preOrderBinaryTree(node *TreeNode) {
	//从root节点开始遍历
	if node != nil {
		b.Callback(node.Key)
		b.preOrderBinaryTree(node.LeftChild)
		b.preOrderBinaryTree(node.RightChild)
	}
}

//遍历二叉树：中序遍历：左->中->右
func (b *BinaryTreeNode) inOrderBinaryTree(node *TreeNode) {
	//从root节点开始遍历
	if node != nil {
		b.inOrderBinaryTree(node.LeftChild)
		b.Callback(node.Key)
		b.inOrderBinaryTree(node.RightChild)
	}
}

//遍历二叉树：后序遍历：左->右->中
func (b *BinaryTreeNode) postOrderBinaryTree(node *TreeNode) {
	//从root节点开始遍历
	if node != nil {
		b.postOrderBinaryTree(node.LeftChild)
		b.postOrderBinaryTree(node.RightChild)
		b.Callback(node.Key)
	}
}

//中序遍历的非递归实现
/**
* 实现原理：利用stack先进后出的特性，
* 第一次遍历，将第一个叶结点的路径节点存入stack，大的先入栈，再将stack里面的元素推出，小的先出来，大的在后面
**/
func (b *BinaryTreeNode) nonRecInOrderBinaryTree() {
	//从root节点开始遍历
	newNode := b.Root
	stack := list.New()
	for newNode != nil || stack.Len() != 0 {
		//第一次遍历，将第一个叶结点的路径节点存入stack，大的先入栈
		for newNode != nil {
			//入栈
			stack.PushBack(newNode)
			newNode = newNode.LeftChild
		}
		//将stack里面的元素推出，小的先出来，大的在后面
		if stack.Len() != 0 {
			//返回最后一个节点
			lastTreeNode := stack.Back()
			newNode = lastTreeNode.Value.(*TreeNode)
			//打印节点
			b.Callback(newNode.Key)
			//遍历当前打印节点的右节点
			newNode = newNode.RightChild
			//移除已打印的节点
			stack.Remove(lastTreeNode)
		}
	}
}

//二叉树深度
func (b *BinaryTreeNode) maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	var leftHeight int = b.maxDepth(node.LeftChild)
	var rightHeight int = b.maxDepth(node.RightChild)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1

}

func main() {
	tree := BinaryTreeNode{
		Root: nil,
		Keys: []int{8, 3, 10, 1, 6, 6, 14, 4, 7, 13},
		Callback: func(key int) {
			fmt.Print(key, "->")
		},
	}
	tree.BuildBinaryTree()
	//从root节点开始遍历：中序遍历
	fmt.Print("\n递归中序遍历:")
	tree.inOrderBinaryTree(tree.Root)

	//中序遍历，非递归方式
	fmt.Print("\n非递归中序遍历:")
	tree.nonRecInOrderBinaryTree()
	//二叉树最大深度
	height := tree.maxDepth(tree.Root)
	fmt.Println("\n最大深度:", height)
}

/**
 *                    [8]
 *               [3]      [10]
 *            [1]   [6]       [14]
 *                [4] [6]   [13]
 *                      [7]
 *
 */
