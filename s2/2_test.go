package s2

import (
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 :=&ListNode{1,nil}
	l1Index :=l1
	l1Index.Next = &ListNode{8,nil}
	l1Index = l1Index.Next
	l1Index.Next = &ListNode{3,nil}


	l2 :=&ListNode{0,nil}
	l2Index :=l2
	l2Index.Next = &ListNode{6,nil}
	l2Index = l2Index.Next
	l2Index.Next = &ListNode{4,nil}
	ret :=AddTwoNumbers(l1,l2)
	retStr :=""
	for ret!=nil {
		retStr+=strconv.Itoa(ret.Val)
		ret= ret.Next
	}
	t.Log(retStr)
}
