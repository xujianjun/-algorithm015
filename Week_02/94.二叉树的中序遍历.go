package Week_02

func inorderTraversal(root *TreeNode) []int {
	return inorderRecursive(root)
}

func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := append(inorderRecursive(root.Left), []int{root.Val}...)
	ret = append(ret, inorderRecursive(root.Right)...)
	return ret
}
