/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValid(root, nil, nil)
}


func isValid(root *TreeNode, floor, ceil *int) bool {
    if root == nil {
        return true
    }
    
    if floor != nil && (*floor >= root.Val) {
        return false
    }
    
    if ceil != nil && (*ceil <= root.Val) {
        return false
    }
    
    if root.Left != nil && root.Left.Val >= root.Val {
        return false
    }
    
    if root.Right != nil && root.Right.Val <= root.Val {
        return false
    }
    
    return isValid(root.Left, floor, min(&root.Val, ceil)) && isValid(root.Right, max(&root.Val, floor), ceil)
}

func min(a, b *int) *int {
    if a == nil {
        return b
    }
    
    if b == nil {
        return a
    }
    
    if *a < *b {
        return a
    }
    
    return b
}

func max(a, b *int) *int {
    if a == nil {
        return b
    }
    
    if b == nil {
        return a
    }
    
    if *a > *b {
        return a
    }
    
    return b
}