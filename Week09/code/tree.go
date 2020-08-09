package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(arr []int, root *TreeNode, i, n int) *TreeNode {
	if i < n {
		temp := TreeNode{
			Val:   arr[i],
			Left:  nil,
			Right: nil,
		}
		root = &temp
		root.Left = CreateTree(arr, root.Left, i*2+1, n)
		root.Right = CreateTree(arr, root.Right, i*2+2, n)
	}

	return root
}
