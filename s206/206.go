package s206

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

思路：
直接遍历一遍。没有什么好招。
*/

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	var ret *ListNode = nil

	for head !=nil{
		tmp := head
		head=head.Next

		tmp.Next = ret
		ret = tmp

	}

	return ret
}
