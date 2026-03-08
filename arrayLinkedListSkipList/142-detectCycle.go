package week01

import "github.com/TylerTang06/-algorithm015/util"

/*
142.环形链表2

给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。
注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	// 先确定有环，A = X(B+C)-B => A = x(B+C)+C
	// b点为第一次相遇的点
	cur, nxt := head.Next, head.Next.Next
	for nxt.Next != nil && nxt.Next.Next != nil {
		if cur == nxt {
			nxt = head
			break
		}
		cur, nxt = cur.Next, nxt.Next.Next
	}

	// a点为第二次相遇，即为环的入口
	for nxt != nil && nxt != cur {
		nxt, cur = nxt.Next, cur.Next
	}

	return nxt
}
