package week01

import "github.com/TylerTang06/-algorithm015/util"

/*
92.反转链表2

给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

示例 1:
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]

提示：
链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *util.ListNode, m int, n int) *util.ListNode {
	if m == n {
		return head
	}

	// notice: bef.Next = head
	curH, bef := head, &util.ListNode{Next: head}
	for i := 0; i < m-1 && curH != nil; i++ {
		curH, bef = curH.Next, bef.Next
	}
	bef.Next = nil // disconnet bef and curH

	cur, curB := curH.Next, curH // curB record the end of curH
	curH.Next = nil              // get the firt node of curH
	for i := m; i < n && cur != nil; i++ {
		nxt := cur.Next
		cur.Next, curH = curH, cur
		cur = nxt
	}

	if m == 1 {
		head, curB.Next = curH, cur
	} else {
		bef.Next, curB.Next = curH, cur
	}

	return head
}
