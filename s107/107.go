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

思路就是深度递归，构成数组，把当前节点加入到结果集
然后翻转数组
翻转数组 只需要遍历到中间就可以停下
因为是将头尾调换
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


	// 翻转数组 只需要遍历整个数组的一半即可
	// 这里注意折半后要减一
	retLen :=len(ret)
	end  :=retLen/2
	for i := 0; i < end; i++ {
		ret[i],ret[retLen-1-i] = ret[retLen-1-i],ret[i]
	}

	return ret
}

func levelOrderBottomLoop(nood *TreeNode,ret *[][]int,index int)  {
	// 只要节点不为nil 就可以添加
	if nood!=nil{

		index ++

		// 处理当前层的子节点
		// 先处理左边
		temp :=[]int{}
		if nood.Left !=nil{
			temp = append(temp, nood.Left.Val)

		}
		// 然后处理右边
		if nood.Right !=nil{
			temp = append(temp, nood.Right.Val)

		}
		// 合并当前的层到ret中
		// temp 大于 0 说明 左右至少有一个值
		if len(temp)>0{
			// 将同一层的节点全部合并
			// 如果ret的长度>index 说明该层已经存在
			// 直接将现在的值append到这一层所属的数组中
			if len(*ret)>index{
				(*ret)[index] = append((*ret)[index], temp...)
			}else{
				// 否则 新建一个层
				*ret = append(*ret, temp)
			}
		}

		levelOrderBottomLoop(nood.Left,ret,index)
		levelOrderBottomLoop(nood.Right,ret,index)

	}
}
