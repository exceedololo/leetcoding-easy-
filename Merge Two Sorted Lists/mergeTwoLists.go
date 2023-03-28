package Merge_Two_Sorted_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sortedList := &ListNode{}
	curr := sortedList
	curr1 := list1
	curr2 := list2

	for curr1 != nil && curr2 != nil {
		if curr1.Val <= curr2.Val {
			curr.Next = &ListNode{Val: curr1.Val}
			curr1 = curr1.Next
		} else {
			curr.Next = &ListNode{Val: curr2.Val}
			curr2 = curr2.Next
		}
		curr = curr.Next
	}
	for curr1 != nil {
		curr.Next = &ListNode{Val: curr1.Val}
		curr1 = curr1.Next
		curr = curr.Next
	}
	for curr2 != nil {
		curr.Next = &ListNode{Val: curr2.Val}
		curr2 = curr2.Next
		curr = curr.Next
	}
	return sortedList.Next
}
