package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	if root.Left == nil && root.Right == nil {
		return true, 1
	}

	checkL, cntL := dfs(root.Left)
	checkR, cntR := dfs(root.Right)
	if !checkL || !checkR {
		return cntL + cntR
	}

	if root.Left != nil && root.Left.Val != root.Val {
		return false, cntL + cntR
	}

	if root.Right != nil && root.Right.Val != root.Val {
		return false, cntL + cntR
	}

	return true, cntL + cntR + 1
}

func countUnivalSubtrees(root *TreeNode) int {
	_, total := dfs(root)
	return total
}

func main() {
}
