package week01

import "github.com/TylerTang06/-algorithm015/util"

/*
206.反转链表

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head.Next
	head.Next = nil
	// head = &ListNode{Next: nil}
	for cur != nil {
		nxt := cur.Next
		cur.Next, head = head, cur
		cur = nxt
	}

	return head
}
