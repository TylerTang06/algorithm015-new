package week04

import (
	"container/list"

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
func levelOrder(root *util.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	myQue := list.New()
	myQue.PushBack(root)
	res := [][]int{}
	for myQue.Len() > 0 {
		l := myQue.Len()
		layer := []int{}
		for l > 0 {
			node := myQue.Front().Value.(*util.TreeNode)
			if node.Left != nil {
				myQue.PushBack(node.Left)
			}
			if node.Right != nil {
				myQue.PushBack(node.Right)
			}
			layer = append(layer, node.Val)
			myQue.Remove(myQue.Front())
			l--
		}
		res = append(res, layer)
	}

	return res
}
