package s100

/*
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true


思路：
同时遍历两个树即可
记住遍历树的要点，以及及时结束条件
*/


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p !=nil && q!=nil{
		if p.Val!=q.Val{
			return false
		}else  {
			left :=isSameTree(p.Left,q.Left)
			right:=isSameTree(p.Right,q.Right)
			if left==true && right ==true{
				return true
			}else {
				return false
			}
		}
	}else {
		if p == q{
			return true
		}
		return false
	}

}
