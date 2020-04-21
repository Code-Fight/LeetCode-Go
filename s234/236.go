package s234

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true

思路：
首先，你要明确什么是回文数？！！
-->链表为空或者只有一个node的时候，都是回文数

然后，判断回文数的办法有两种，一个是双指针，一头一尾进行对比，如果条件不够，那就只能先翻转一下给定的链表，然后 对比两个链表。

*/


type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head ==nil || head.Next==nil{
		return true
	}

	var flip *ListNode = nil

	var temp *ListNode = head
	for temp!=nil{
		t := *temp
		temp = temp.Next
		t.Next = flip
		flip = &t
	}


	for flip!=nil && head!=nil {
		if flip.Val != head.Val{
			return false
		}else {
			flip = flip.Next
			head = head.Next
		}
	}

	if flip!=nil || head !=nil{
		return false
	}
	return true

}
