package Balanced_Binary_Tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	var leftHeight, rightHeight int
	if isBalanced(root.Left) {
		leftHeight = getHeight(root.Left)
	} else {
		return false
	}
	if isBalanced(root.Right) {
		rightHeight = getHeight(root.Right)
	} else {
		return false
	}
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}
	return true
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if rightHeight > leftHeight {
		return rightHeight + 1
	}
	return leftHeight + 1
}
