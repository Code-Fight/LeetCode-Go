package s257

import "testing"

func TestBinaryTreePathsTest(t *testing.T) {
	root:=TreeNode{Val: -1,Left: &TreeNode{Val: 2,Left: nil,Right: &TreeNode{Val: 5}},Right: &TreeNode{Val: 3}}
	binaryTreePaths(&root)
}