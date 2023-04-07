package Intersection_of_Two_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curr1 := headA
	curr2 := headB
	if curr1 == nil || curr2 == nil {
		return nil
	}
	for curr1 != curr2 {
		if curr1 == nil {
			curr1 = headB
		} else {
			curr1 = curr1.Next
		}
		if curr2 == nil {
			curr2 = headA
		} else {
			curr2 = curr2.Next
		}
	}
	return curr1
}
