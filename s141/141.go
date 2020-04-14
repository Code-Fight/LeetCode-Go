package s141


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
		slowIndex = slowIndex.Next
		if slowIndex ==nil{
			return false
		}
		fastIndex = fastIndex.Next
		if fastIndex ==nil {
			return false
		}
		fastIndex = fastIndex.Next
		if fastIndex ==nil {
			return false
		}

		if slowIndex == fastIndex{
			return true
		}

	}

	return false
}
