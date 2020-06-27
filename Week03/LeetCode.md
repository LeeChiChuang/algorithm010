## #22 [括号生成](https://leetcode-cn.com/problems/generate-parentheses/submissions/)

```go
func generateParenthesis(n int) []string {
    res := new([]string)
    generate(0, 0, n, "", res)
    return *res
}

func generate(left int, right int, n int, str string, res *[]string) {
    if left == n && right == n {
        *res = append(*res, str)
        return
    }

    if left < n {
        generate(left+1, right, n, str+"(", res)
    }

    if right < left {
        generate(left, right+1, n, str+")", res)
    }
}
```



##  105 [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{preorder[0], nil, nil}
    i := 0
    for ; i<len(inorder); i++ {
        if preorder[0] == inorder[i] {
            break
        }
    }

    root.Left = buildTree(preorder[1:i+1], inorder[:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])
    return root

```



## 236 [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root == nil {
         return nil
     }
     if root.Val == p.Val || root.Val == q.Val {
         return root
     }
     left := lowestCommonAncestor(root.Left, p, q)
     right := lowestCommonAncestor(root.Right, p, q)
     if left != nil && right != nil {
         return root
     }
     if left == nil {
         return right
     }
     return left
}
```



## 77 [组合](https://leetcode-cn.com/problems/combinations/)

```go
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, k)
	var dfs func(start int, idx int)
	dfs = func(start int, idx int) {
		for i := start; i <= n-(k-1-idx); i++ {
			option[idx] = i
			if idx == k-1 {
				ans = append(ans, append([]int{}, option...))
			} else {
				dfs(i+1, idx+1)
			}
		}
	}
	dfs(1, 0)
	return ans
}
```



## 46 [全排列](https://leetcode-cn.com/problems/permutations/)



```go
var result [][]int

func permute(nums []int) [][]int {
    result = [][]int{}
	appendPath := make([]int, 0)
	used := make(map[int]bool, len(nums))
	backTrace(nums, appendPath, used)
	return result
}

func backTrace(nums []int, appendPath []int, used map[int]bool) {
	if len(nums) == len(appendPath) {
		tmp := make([]int, len(nums))
		copy(tmp, appendPath)
		result = append(result, tmp)
		return
	}

	for i:=0; i<len(nums); i++ {
		if !used[i] {
			used[i] = true
			appendPath = append(appendPath, nums[i])
			backTrace(nums, appendPath, used)
			appendPath = appendPath[:len(appendPath)-1]
			used[i] = false
		}
	}
}
```

