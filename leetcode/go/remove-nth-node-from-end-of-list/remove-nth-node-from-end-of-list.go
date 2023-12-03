/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	l := length(head)

	return helper(head, 0, l-n)
}

func helper(head *ListNode, idx, removeIdx int) *ListNode {
	if head == nil {
		return nil
	}

	if removeIdx == 0 {
		return head.Next
	}

	if removeIdx == idx {
		return head.Next
	}

	head.Next = helper(head.Next, idx+1, removeIdx)

	return head
}

func length(head *ListNode) int {
	if head == nil {
		return 0
	}

	return length(head.Next) + 1
}