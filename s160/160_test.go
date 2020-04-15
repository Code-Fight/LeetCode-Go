package s160

import "testing"

func TestGetIntersectionNode(t *testing.T)  {

	A :=&ListNode{4,nil}
	B :=&ListNode{2,A}


	headA:=&ListNode{0,&ListNode{9,&ListNode{1,nil}}}
	headB:=B

	c:=getIntersectionNode(headA,headB)
	t.Log(c)
}