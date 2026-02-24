package week01

import "github.com/TylerTang06/-algorithm015/util"

/*
141.环形链表
给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。
仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *util.ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	cur, nxt := head.Next, head.Next.Next
	for nxt.Next != nil && nxt.Next.Next != nil {
		if cur == nxt {
			return true
		}
		cur, nxt = cur.Next, nxt.Next.Next
	}

	return false
}
