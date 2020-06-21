

学习笔记

前序遍历(Pre-order)： 根-左-右

中序遍历(In-order): 左根右

后序遍历(Post-order): 左右根

## hash

### Map 

hashMap个人总结

### Set



## 实战题目

### [有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/description/)（亚马逊、Facebook、谷歌在半年内面试中考过）

> 字母次数一样 顺序不一样
>
> clarification
>
> possible solution >> time & space
>
> code
>
> test case
>
> 1. 暴力 sort >> 相等 O(NlogN)



```go
import (
    "strings"
    "sort"
)
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sSlice := strings.Split(s, "")
    tSlice := strings.Split(t, "")
    sort.Strings(sSlice)
    sort.Strings(tSlice)

    for index, ele := range sSlice {
        if ele != tSlice[index] {
            return false
        }
    }

    return true
}
```



```go
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    arrS := [26]int{}
    arrT := [26]int{}

    for _, v := range(s) {
        arrS[v - 'a'] += 1
    }
    for _, v := range(t) {
        arrT[v - 'a'] += 1
    }

    return arrS == arrT
}
```



```go

```



### [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)（亚马逊在半年内面试中常考）



### [两数之和](https://leetcode-cn.com/problems/two-sum/description/)（亚马逊、字节跳动、谷歌、Facebook、苹果、微软、腾讯在半年内面试中常考）



### [二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)（亚马逊、微软、字节跳动在半年内面试中考过）

> 递归

```go
var traversePath []int
func inorderTraversal(root *TreeNode) []int {
    traversePath = make([]int, 0)
    return dfs(root)   
}

func dfs(root *TreeNode) []int {
    if root != nil {
        dfs(root.Left)
        traversePath = append(traversePath, root.Val)
        dfs(root.Right)
    }

    return traversePath
}
```

> **Stack**迭代 时间复杂度O(n) 空间复杂度O(n)

```go
func inorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var res []int
    index := 0
    
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        index = len(stack) - 1
        res = append(res, stack[index].Val)
        root = stack[index].Right
        stack = stack[:index]
    }

    return res
}
```

>莫里斯遍历 **破坏树结构**
>
>curr初始化为root
>
>当curr >> left 为空
>
>​	append元素
>
>​	curr >> right
>
>不为空
>
>​	prev >> curr.left 然后找到prev的最右子节点
>
>​	pref >> curr
>
>​	curr >> curr.left

```go
func inorderTraversal(root *TreeNode) []int {
    var res []int
    var prev *TreeNode
    curr := root

    for curr != nil {
        if curr.Left == nil {
            res = append(res, curr.Val)
            curr = curr.Right
        } else {
            prev = curr.Left
            for prev.Right != nil {
                prev = prev.Right
            }
            prev.Right = curr
            tmp := curr
            curr = curr.Left
            tmp.Left = nil
        }
    }

    return res
}
```



### [二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)（谷歌、微软、字节跳动在半年内面试中考过）

### [N 叉树的后序遍历](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)（亚马逊在半年内面试中考过）

### [N 叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/)（亚马逊在半年内面试中考过）

>递归

```go
var res []int
func preorder(root *Node) []int {
    res = make([]int , 0)
    if root == nil {
        return res
    }
    res = append(res, root.Val)
    return recursive(root.Children)
}

func recursive(children []*Node) []int {
    if children != nil {
        for _, v := range(children) {
            res = append(res, v.Val)
            recursive(v.Children)
        }
    }

    return res
}
```



### [N 叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

```go
func levelOrder(root *Node) [][]int {
    if root == nil {
        return [][]int{}
    }
    quene := make([]*Node, 0)
    quene = append(quene, root)
    res := make([][]int, 0)
    for len(quene) != 0 {
        level := make([]int, 0)
        size := len(quene)
        for i:=0; i<size; i++ {
            level = append(level, quene[0].Val)
            quene  = append(quene, quene[0].Children...)
            quene = quene[1:]
        }
        res = append(res, level)
    }
    return res
}
```





## **中等：**

## [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)（亚马逊在半年内面试中常考）

## [二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)（亚马逊、字节跳动、微软在半年内面试中考过）

## [二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)（字节跳动、谷歌、腾讯在半年内面试中考过）

## [N 叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)（亚马逊在半年内面试中考过）

## [丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)（字节跳动在半年内面试中考过）

```go
func nthUglyNumber(n int) int {
    // a*2 b*3 c*5
    a, b, c := 0,0,0 // index
    uglyNums := make([]int, n)
    uglyNums[0] = 1
    for i:=1;i<n;i++ {
        numA, numB, numC := uglyNums[a]*2, uglyNums[b]*3, uglyNums[c]*5
        minVal := min(numA, numB, numC)
        if minVal == numA {
            a++
        }
        if minVal == numB {
            b++
        }
        if minVal == numC {
            c++
        }
        uglyNums[i] = minVal
    }
    //fmt.Println(uglyNums)
    return uglyNums[n-1]
}

func min(x int, y int, z int) int {
    if x <= y && x <= z {
        return x
    }

    if y <= x && y <= z {
        return y
    }

    return z
}
```



## [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)（亚马逊在半年内面试中常考）