/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil && l1 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		l2 = &ListNode{}
	}

	sum := l1.Val + l2.Val

	remain := sum / 10

	if remain > 0 {
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}

		l1.Next.Val += remain
	}

	return &ListNode{
		Val:  sum % 10,
		Next: addTwoNumbers(l1.Next, l2.Next),
	}
}
