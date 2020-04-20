package s206

import "testing"

func TestReverseList(t *testing.T) {
	temp:=ListNode{0,&ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,nil}}}}}
	reverseList(&temp)
}
