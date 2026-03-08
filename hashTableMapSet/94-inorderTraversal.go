package week02

import (
	"container/list"

	"github.com/TylerTang06/-algorithm015/util"
)

/*
94.二叉树中序遍历

给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *util.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	myStack := list.New()
	myMap := map[*util.TreeNode]int{}

	myStack.PushBack(root)
	myMap[root] = 1
	for myStack.Len() > 0 {
		node := myStack.Back().Value.(*util.TreeNode)
		myStack.Remove(myStack.Back())
		if 1 == myMap[node] {
			if node.Right != nil {
				myStack.PushBack(node.Right)
				myMap[node.Right] = 1
			}
			myStack.PushBack(node)
			myMap[node] = 2
			if node.Left != nil {
				myStack.PushBack(node.Left)
				myMap[node.Left] = 1
			}
		} else {
			res = append(res, node.Val)
		}
	}

	return res
}

func inorderTraversal1(root *util.TreeNode) []int {
	var res []int
	var inOrder func(node *util.TreeNode)
	inOrder = func(node *util.TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}

	inOrder(root)

	return res
}

func inorderTraversal2(root *util.TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			// 思路：中序遍历访：root的前一个访问节点是左子树的最右子树节点，
			// 借助辅助索引将左子树的最右节点的右子树指针链接到root节点
			predecessor := root.Left
			// 获取左子树的最右子树节点指针
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				root = root.Left
				predecessor.Right = root
			} else {
				// 左子树的所有节点已经访问完成后，记录root，并断开辅助索引，开始访问root右子树
				res = append(res, root.Val)
				predecessor.Right = nil
				root = root.Right
			}
		}
	}

	return res
}
