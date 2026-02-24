package week03

import "github.com/TylerTang06/-algorithm015/util"

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
