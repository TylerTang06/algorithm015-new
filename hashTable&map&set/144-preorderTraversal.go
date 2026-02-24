package week02

import (
	"container/list"

	"github.com/TylerTang06/-algorithm015/util"
)

/*
144.二叉树的前序遍历
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *util.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	myStack := list.New()
	myStack.PushBack(root)
	for myStack.Len() > 0 {
		node := myStack.Back().Value.(*util.TreeNode)
		myStack.Remove(myStack.Back())
		res = append(res, node.Val)
		if node.Right != nil {
			myStack.PushBack(node.Right)
		}
		if node.Left != nil {
			myStack.PushBack(node.Left)
		}
	}

	return res
}

func preorderTraversal1(root *util.TreeNode) []int {
	var res []int
	var preOder func(node *util.TreeNode)
	preOder = func(node *util.TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		preOder(node.Left)
		preOder(node.Right)
	}

	preOder(root)

	return res
}

func preorderTraversal2(root *util.TreeNode) []int {
	var res []int
	var p1, p2 *util.TreeNode = root, nil
	for p1 != nil {
		// 思路：使用辅助指针降低空间复杂度
		p2 = p1.Left
		// 如果左子树为空，则记录节点
		if p2 == nil {
			res = append(res, p1.Val)
		} else {
			// 获取左子树的最右子树指针
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}

			// 如果最右指针的右子树为空，则记录节点，并将最右指针指向节点并继续访问节点左子树
			if p2.Right == nil {
				res = append(res, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			// 如果最右指针为右子树为空，说明已经记录了节点，断开辅助索引
			p2.Right = nil
		}

		// 继续访问节点右子树
		p1 = p1.Right

	}

	return res
}
