package more

import (
	"math"

	"github.com/TylerTang06/-algorithm015/util"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *util.TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return int(^uint(0) >> 1)
	}
	l, r := int(^uint(0)>>1), int(^uint(0)>>1)
	if root.Left != nil {
		l = maxNumInBST(root.Left)
		l = int(math.Abs(float64(root.Val - l)))
	}
	if root.Right != nil {
		r = minNumInBST(root.Right)
		r = int(math.Abs(float64(root.Val - r)))
	}
	// fmt.Println(l, r)
	if l > r {
		l = r
	}
	ld := minDiffInBST(root.Left)
	rd := minDiffInBST(root.Right)
	// fmt.Println(l, ld, rd)
	if l > ld {
		l = ld
	}
	if l > rd {
		return rd
	}
	return l
}

func maxNumInBST(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right != nil {
		return maxNumInBST(root.Right)
	}
	return root.Val
}

func minNumInBST(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		return minNumInBST(root.Left)
	}
	return root.Val
}
