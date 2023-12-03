/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    return bruchforce(head, n)
}

func bruchforce(head *ListNode, n int) *ListNode {
    l := length(head)
    pos := l - n 
    cur := head
    prev := &ListNode{}

    for i := 0; i <= pos; i++ {
        if pos == 0 {
            return head.Next
        }
        
        if i == pos {
            prev.Next = cur.Next
            return head
        }
        
        prev = cur
        cur = cur.Next
    }
    
    return prev
}


func length(head *ListNode) int {
    if head == nil {
        return 0
    }
    
    return 1 + length(head.Next)
}