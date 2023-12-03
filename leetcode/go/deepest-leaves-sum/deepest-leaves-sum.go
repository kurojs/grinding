/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    sum, _ := sumLeaves(root, 0)
    return sum
}

func sumLeaves(root *TreeNode, level int) (int, int) {
    if root == nil {
        return 0, level
    }
    
    leftSum, leftLevel := sumLeaves(root.Left, level+1)
    rightSum, rightLevel := sumLeaves(root.Right, level+1)
    
    if leftLevel == rightLevel {
        if leftSum == 0 {
            return root.Val, level+1
        }
        
        return leftSum + rightSum, leftLevel
    }
    
    if leftLevel > rightLevel {
        return leftSum, leftLevel
    }
    
    return rightSum, rightLevel
}


//  root 