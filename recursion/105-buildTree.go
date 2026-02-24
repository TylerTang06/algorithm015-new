package week03

import "github.com/TylerTang06/-algorithm015/util"

/*
105.从前序和中序遍历序列构建二叉树
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *util.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := &util.TreeNode{Val: preorder[0], Left: nil, Right: nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	if i == len(inorder) {
		return nil
	}
	rootVal.Left = buildTree(preorder[1:i+1], inorder[:i])
	rootVal.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return rootVal
}
