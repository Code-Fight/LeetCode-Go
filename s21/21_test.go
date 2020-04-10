package s21

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T)  {
	l1 :=&ListNode{1,nil}
	//l1Index :=l1
	//l1Index.Next = &ListNode{2,nil}
	//l1Index = l1Index.Next
	//l1Index.Next = &ListNode{4,nil}

	var l2 *ListNode
	//l2 :=&ListNode{1,nil}
	//l2Index :=l2
	//l2Index.Next = &ListNode{3,nil}
	//l2Index = l2Index.Next
	//l2Index.Next = &ListNode{4,nil}
	ret:=mergeTwoLists(l1,l2)
	for ret!=nil {
		t.Log(ret.Val,"-")
		ret=ret.Next
	}
}
