package s234

import "testing"

func TestIsPalindrome(t *testing.T) {
	temp:=ListNode{0,&ListNode{1,&ListNode{2,&ListNode{1,&ListNode{0,nil}}}}}
	isPalindrome(&temp)
}
