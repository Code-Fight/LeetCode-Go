package s104

import "testing"

func TestMaxDepth(t *testing.T) {

	t3 := &TreeNode{0,nil,nil}

	t2 := &TreeNode{0,nil,nil}
	t2_1 := &TreeNode{0,t3,nil}
	t1 := &TreeNode{0,nil,nil}
	t1_1 := &TreeNode{0,t2,t2_1}

	t0 :=TreeNode{0,t1,t1_1}

	depth:=maxDepth(&t0)

	t.Log(depth==4)
}
