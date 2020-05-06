package s107

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

思路：
这个同一层的都属于一个数组内，别搞混了

思路就是深度递归，构成数组，然后翻转数组
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret :=[][]int{}
	if root==nil{
		return ret
	}
	ret = append(ret,[]int{root.Val} )
	levelOrderBottomLoop(root,&ret,0)


	retLen :=len(ret)
	end  :=retLen/2
	for i := 0; i < end; i++ {
		ret[i],ret[retLen-1-i] = ret[retLen-1-i],ret[i]
	}

	return ret
}

func levelOrderBottomLoop(nood *TreeNode,ret *[][]int,index int)  {

	if nood!=nil{

		index ++

		temp :=[]int{}
		if nood.Left !=nil{
			temp = append(temp, nood.Left.Val)

		}
		if nood.Right !=nil{
			temp = append(temp, nood.Right.Val)

		}
		if len(temp)>0{
			if len(*ret)>index{
				(*ret)[index] = append((*ret)[index], temp...)
			}else{
				*ret = append(*ret, temp)
			}
		}

		levelOrderBottomLoop(nood.Left,ret,index)
		levelOrderBottomLoop(nood.Right,ret,index)

	}
}
