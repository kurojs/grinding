/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    m := make(map[int][]int)
    helper(root, 0, m)
    
    level := len(m)
    
    ans := make([][]int, level)
    for i := 0; i < level; i++ {
        ans[i] = m[i]
    }
    
    return ans
}

func helper(root *TreeNode, level int, m map[int][]int) {
    if root == nil {
        return 
    }
    
    if _, ok := m[level]; !ok {
        m[level] = make([]int, 0)
    }
    
    m[level] = append(m[level], root.Val)
    
    helper(root.Left, level + 1, m)
    helper(root.Right, level + 1, m)
}