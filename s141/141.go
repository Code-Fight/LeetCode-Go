package s141

/*
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。

 

进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

思路：
判断链表是否有环，直接快慢指针，如果有换必定相交
如果，知道遍历结束都没有相交，说明没有环
*/


type ListNode struct {
    Val int
    Next *ListNode
}
 
func hasCycle(head *ListNode) bool {
	if head ==nil{
		return false
	}

	slowIndex:=head
	fastIndex:=head

	for slowIndex!=nil&&fastIndex!=nil{
		// 慢指针前进一步
		slowIndex = slowIndex.Next
		if slowIndex ==nil{
			return false
		}

		// 快指针前进两步，但是需要注意第二次前进的时候，先判断是否有下一个节点
		fastIndex = fastIndex.Next
		if fastIndex ==nil {
			return false
		}
		fastIndex = fastIndex.Next
		if fastIndex ==nil {
			return false
		}

		// 如果 快慢相等 说明有环了
		if slowIndex == fastIndex{
			return true
		}

	}

	return false
}
