package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	level := BTreeLevelCount(root)
	for i := 1; i <= level; i++ {
		printLevel(root, i, f)
	}
}

func printLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		printLevel(root.Left, level-1, f)
		printLevel(root.Right, level-1, f)
	}
}
