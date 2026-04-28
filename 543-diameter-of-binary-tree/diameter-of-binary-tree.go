/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxDiametr int

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    maxDiametr = 0
    maxDepth(root)
    
    return maxDiametr
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    
    currentDiametr := left + right

    if currentDiametr > maxDiametr {
        maxDiametr = currentDiametr
    }

    return max(right, left) + 1
}