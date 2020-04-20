package s206


type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	var ret *ListNode = nil

	for head !=nil{
		tmp := head
		head=head.Next

		tmp.Next = ret
		ret = tmp

	}

	return ret
}
