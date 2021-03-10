package s257

/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

思路：深度遍历，然后注意控制nil，没深入一次，添加当前节点到当前的path中，
     如果lef和right都为nil的时候，说明为最后的节点了，然后就加入到返回结果paths中
*/

import (
	"strconv"
	"strings"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	// 最后的返回值
	paths :=[]string{}
	// 从跟节点到叶子节点的路径
	path := ""
	if root !=nil{
		loop(root,path,&paths)
	}

	return paths
}

func loop(root *TreeNode,path string,paths *[]string)  {


	if len(path)>0{
		path += "->"+ strconv.Itoa(root.Val)
	}else{
		path += strconv.Itoa(root.Val)
	}

	if root.Left!=nil{
		loop(root.Left,path,paths)
	}
	//temp =*path + "->"+ strconv.Itoa(root.Val)
	if root.Right!=nil{
		loop(root.Right,path,paths)
	}

	// 如果 左右都为nil 说明目前为叶子节点
	// 那么 就需要把当前的path 加入到paths中
	if root.Left==nil && root.Right==nil{
		*paths = append(*paths, strings.TrimPrefix(path,"->"))
	}
}


