/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// Pre proccessing index mapping
	mPre, mIn := make(map[int]int), make(map[int]int)
	for i, num := range preorder {
		mPre[num] = i
	}
	for i, num := range inorder {
		mIn[num] = i
	}

	return build(preorder, 0, len(preorder)-1, mPre, inorder, 0, len(inorder)-1, mIn)
}


func build(preorder []int, prestart, preend int, mPre map[int]int, inorder []int, instart, inend int, mIn map[int]int) *TreeNode {
	if prestart > preend || instart > inend {
		return nil
	}
    
	root := &TreeNode{
		Val: preorder[prestart],
	}

	/*
    (a) Inorder (Left, Root, Right) : 9 3 15 20 7
    (b) Preorder (Root, Left, Right): 3 9 20 15 7
    */
    
	inRootIdx := mIn[root.Val]
    inLeftNodes := inRootIdx - instart
    
	root.Left = build(preorder, prestart + 1, prestart + inLeftNodes, mPre, inorder, instart, inRootIdx - 1, mIn)
	root.Right = build(preorder, prestart + inLeftNodes + 1, preend, mPre, inorder, inRootIdx+1, inend, mIn)

	return root
}
