package s160


/*
 相交链表
编写一个程序，找到两个单链表相交的起始节点。

思路：
1.简单的办法就是 做一个hashtable 然后遍历之后判断是否存在相同的

2.还有一种思路利用现有的利用双指针，如果链表相交，那么肯定两个指针会相同。
创建两个指针 pApA 和 pBpB，分别初始化为链表 A 和 B 的头结点。然后让它们向后逐结点遍历。
当 pApA 到达链表的尾部时，将它重定位到链表 B 的头结点 (你没看错，就是链表 B); 类似的，当 pBpB 到达链表的尾部时，将它重定位到链表 A 的头结点。
若在某一时刻 pApA 和 pBpB 相遇，则 pApA/pBpB 为相交结点。
想弄清楚为什么这样可行, 可以考虑以下两个链表: A={1,3,5,7,9,11} 和 B={2,4,9,11}，相交于结点 9。
由于 B.length (=4) < A.length (=6)，pBpB 比 pApA 少经过 2 个结点，会先到达尾部。
将 pBpB 重定向到 A 的头结点，pApA 重定向到 B 的头结点后，pBpB 要比 pApA 多走 2 个结点。因此，它们会同时到达交点。
如果两个链表存在相交，它们末尾的结点必然相同。因此当 pApA/pBpB 到达链表结尾时，记录下链表 A/B 对应的元素。若最后元素不相同，则两个链表不相交。


所以可以参考下面的代码
注意点：为什么指针从A切换到B上之后，没有造成死循环，因为，两个链表都要走一遍，那么总路程相同，终点都是nil，所以会在nil的时候 停止循环

*/

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA := headA
	currB := headB

	// 这里这个判断是防止死循环的关键点
	// 因为涉及到指针切换链表，两切两个指针最后要么到达相交的点，
	//  要么同时到达两个链表的nil  这个同时到达nil 就会停止循环 是一个关键点
	for currA !=currB{
		if currA!=nil{
			currA = currA.Next
		}else {
			currA = headB
		}

		if currB!=nil{
			currB = currB.Next
		}else {
			currB = headA
		}

	}
	return currA
}


// hashtable 方法
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA ==nil || headB ==nil{
//		return nil
//	}
//	headAm := make(map[*ListNode]*ListNode)
//	headBm := make(map[*ListNode]*ListNode)
//
//
//	for headA !=nil || headB!=nil{
//		if headA == headB{
//			return headA
//		}
//
//		if headB!=nil&& headAm[headB] == headB{
//			return headB
//		}
//
//		if headA!=nil &&headBm[headA] == headA{
//			return headA
//		}
//
//		headAm[headA] = headA
//		headBm[headB] = headB
//
//
//		if headA !=nil{
//			headA = headA.Next
//
//		}
//
//		if headB!=nil{
//			headB = headB.Next
//
//		}
//	}
//
//	return nil
//
//}
