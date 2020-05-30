package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeLevelCount(root *TreeNode) int {
	leftCount := 1
	rightCount := 1
	if root == nil {
		return 0
	}
	if root.Left != nil {
		leftCount = 1 + BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		rightCount = 1 + BTreeLevelCount(root.Right)
	}
	if leftCount > rightCount {
		return leftCount
	} else {
		return rightCount
	}
}
