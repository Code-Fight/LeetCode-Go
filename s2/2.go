package s2

/*
 两数相加
给出两个非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

思路：
手动实现按位的加法，遍历整个链表

*/

type ListNode struct {
	 Val int
	 Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret:=&ListNode{0,nil}
	// 当前为的计算结果
	js:=ret
	//代表进位
	jw:=0
	// 两个链表同时往下走
	// 计算和 计算 进位 计算进位后剩下的值
	for l1!=nil  || l2!=nil {

		l1Int := 0
		if l1 == nil {
			l1Int = 0
		} else {
			l1Int = l1.Val
		}
		l2Int := 0
		if l2 == nil {
			l2Int = 0
		} else {
			l2Int = l2.Val
		}
		js.Val = l1Int + l2Int + jw
		jw = js.Val / 10
		js.Val %= 10

		if l1 !=nil{
			l1 = l1.Next
		}
		if l2!=nil{
			l2 = l2.Next
		}

		if l1!=nil  || l2!=nil{
			js.Next = &ListNode{0,nil}
		}else {
			if jw>0{
				js.Next = &ListNode{jw,nil}
			}
		}

		js = js.Next
	}
	return ret
}
