package Week_02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return preorderRecursive(root)
}

func preorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := append([]int{root.Val}, preorderRecursive(root.Left)...)
	ret = append(ret, preorderRecursive(root.Right)...)
	return ret
}
