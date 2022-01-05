package main

import (
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

//遍历二叉树
func (b *BinaryTreeNode) inOrderBinaryTree(node *TreeNode) {
	//从root节点开始遍历
	if node != nil {
		b.inOrderBinaryTree(node.LeftChild)
		b.Callback(node.Key)
		b.inOrderBinaryTree(node.RightChild)
	}
}

//二叉树深度
func (b *BinaryTreeNode) maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	} else {
		var leftHeight int = b.maxDepth(node.LeftChild)
		var rightHeight int = b.maxDepth(node.RightChild)
		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
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
	//从root节点开始遍历
	tree.inOrderBinaryTree(tree.Root)
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
