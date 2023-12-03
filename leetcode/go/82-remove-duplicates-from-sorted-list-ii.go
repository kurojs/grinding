/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    head = &ListNode{
        Val: -1997,
        Next: head,
    }
    cur := head
    prev := head
    
    for ;cur != nil; {
        next := cur.Next
        isDup := false
        
        for ;next != nil && next.Val == cur.Val; {
            next = next.Next
            isDup = true
        }
        
        if isDup {
            prev.Next = next
            cur = next
        } else {
            prev = cur
            cur = cur.Next            
        }
    }
    
    if head.Val == -1997 {
        return head.Next
    }
    
    return head
}