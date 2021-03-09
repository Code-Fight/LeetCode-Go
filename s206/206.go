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
	// 开辟一个地址来存储翻转后的链表
	var ret *ListNode = nil

	for head !=nil{
		// 把当前的节点 取出来
		tmp := head
		// 并把head前进到下一个节点
		head=head.Next
		// 把当前next 指向ret
		tmp.Next = ret
		// 再把tmp 变为 ret 即可保持ret是刚加入的节点
		ret = tmp

	}

	return ret
}
