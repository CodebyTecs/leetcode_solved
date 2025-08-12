/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    stack := []*TreeNode{}
    stack = append(stack, root)

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        res = append(res, node.Val)


        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }

    for i := 0; i < len(res)/2; i++ {
        res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
    }

    return res
}