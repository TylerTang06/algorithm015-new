package week03

import (
	"github.com/TylerTang06/-algorithm015/util"
)

/*
236.二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *util.TreeNode) *util.TreeNode {
	// terminator
	if root == nil || p == root || q == root {
		return root
	}

	// 左子树查找
	left := lowestCommonAncestor(root.Left, p, q)
	// 右子树查找
	right := lowestCommonAncestor(root.Right, p, q)

	// 1.左子树均无法查找到p，q 则均在右子树
	if left == nil {
		return right
	}
	// 2.右子树均无法查找到p，q 则均在左子树
	if right == nil {
		return left
	}

	// 3.左右子树各一个
	return root
}
