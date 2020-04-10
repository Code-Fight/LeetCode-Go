package s21
/*
题目：
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

思路：
这个题目，从时间复杂度上来讲，至少需要遍历一次，所以没得优化了，只能在遍历的过程中，尽早结束
空间上，不要去开辟一条新的链表空间（那也用了一个节点），用指针的方式，在两条链表之间游走即可
*/



type ListNode struct {
     Val int
     Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

     if l1 ==nil && l2==nil{
          return l1
     }
     if l1==nil{
         return l2
     }
     if l2 ==nil{
          return l1
     }

     var pre *ListNode
     if l1.Val<=l2.Val{
         pre =l1
         l1 = l1.Next
     }else {
         pre = l2
         l2 = l2.Next
     }
     ret :=pre
     for l1!=nil||l2!=nil{
          if l1 ==nil{
               pre.Next = l2
               break
          }
          if l2 ==nil{
               pre.Next =l1
               break
          }
          if l1.Val<=l2.Val{
               pre.Next =l1
               l1 = l1.Next
          }else {
               pre.Next =l2
               l2 = l2.Next
          }
          pre = pre.Next
     }

     return ret

}
