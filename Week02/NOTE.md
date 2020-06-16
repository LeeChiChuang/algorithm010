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

### [二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)（谷歌、微软、字节跳动在半年内面试中考过）

### [N 叉树的后序遍历](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)（亚马逊在半年内面试中考过）

### [N 叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/)（亚马逊在半年内面试中考过）

### [N 叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)