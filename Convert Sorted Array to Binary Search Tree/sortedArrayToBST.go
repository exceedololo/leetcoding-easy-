package Convert_Sorted_Array_to_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	node := new(TreeNode)
	node.Val = nums[middle]
	node.Left = sortedArrayToBST(nums[:middle])
	node.Right = sortedArrayToBST(nums[middle+1:])
	return node
}
