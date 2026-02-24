package week02

import (
	"github.com/TylerTang06/-algorithm015/util"
)

/*
589.n叉树的前序遍历
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *util.Node) []int {
	if root == nil {
		return []int{}
	}

	res := []int{root.Val}
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}

	return res
}
