/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
//var nexta *ListNode
	curr := head
	for curr != nil{
		if curr.Next != nil && curr.Val == curr.Next.Val{
			curr.Next = curr.Next.Next
		}else{
			curr = curr.Next
		}
	}
	return head
}