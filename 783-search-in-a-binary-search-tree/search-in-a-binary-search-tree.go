/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    var result *TreeNode
    var inorder func(node *TreeNode)

    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Val == val {
            result = &TreeNode{
                Val: val,
                Left: node.Left,
                Right: node.Right,
            }
        }
        inorder(node.Left)
        inorder(node.Right)
    }

    inorder(root)
    return result
}