package livetest

func maxDepth(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > res {
			res = depth
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	return res
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var ans bool
	if p == nil && q == nil {
		ans = true
	} else if p == nil || q == nil {
		ans = false
	} else if p.Val == q.Val {
		ans = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		ans = false
	}
	return ans
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		} else if p == nil || q == nil {
			return false
		} else if p.Val != q.Val {
			return false
		}
		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}

	return dfs(root.Left, root.Right)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
