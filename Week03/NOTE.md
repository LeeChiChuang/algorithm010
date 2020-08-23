# 学习笔记



## 递归

- 不要人肉递归
- 找到最近最简单方法，将其拆解成可重复解决的问题（重复子问题）
- 数学归纳法思维

> python代码模板

```python
def recurion(level, param1, param2, ...):
  # recusion terminator
  if level > MAX_LEVEL:
    process_result
    return
  
  # process logic in current level
  process(level, data...)
  
  # drill down
  self.recursion(level+1, p1, ...)
  
 # revers the current level status if needed
```



## 分治 Divide & Conquer

> 分治代码模板

```python
def divide_conquer(proble, param1, param2, ...):
  # recurion terminator
  if problem is None:
    print_result
    return
  
  # prepare data
  data = prepare_data(problem)
  subproblems = split_problem(problem, data)
  
  # conquer subproblems 
  subresult1 = self.divide_conquer(subproblems[0], p1, ...)
  subresult2 = self.divide_conquer(subproblems[1], p1, ...)
  subresult3 = self.divide_conquer(subproblems[2], p1,...)
  ...
  # process and generate the final result
  result = process_result(subresult1, subresult2, subresult3, ...)
  
  # revert the current level states
```



## 回溯

>采用试错的思想，尝试分步解决一个问题。在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，他将取消上一步甚至是上几步的计算，再通过其它的可能分步解答再次尝试寻找答案。

回溯法通常用最简单的递归方法来实现，在反复上述的步骤后可能出现两种情况：

- 找到可能存在的正确答案
- 在尝试了所有的分步方法后宣告该问题没有答案

最坏情况下，回溯法会导致一次复杂度为指数时间的计算



## 泛型递归 树的递归

## 实战题目

- [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)（阿里巴巴、腾讯、字节跳动在半年内面试常考）
- [括号生成](https://leetcode-cn.com/problems/generate-parentheses/) (字节跳动在半年内面试中考过)
- [翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/description/) (谷歌、字节跳动、Facebook 在半年内面试中考过)
- [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree)（亚马逊、微软、Facebook 在半年内面试中考过）
- [二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree)（亚马逊、微软、字节跳动在半年内面试中考过）
- [二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree)（Facebook、字节跳动、谷歌在半年内面试中考过）
- [二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)（Facebook、亚马逊在半年内面试常考）

## 每日一课

- [如何优雅地计算斐波那契数列](https://time.geekbang.org/dailylesson/detail/100028406)

## 课后作业

- [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)（Facebook 在半年内面试常考）
- [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal)（字节跳动、亚马逊、微软在半年内面试中考过）
- [组合](https://leetcode-cn.com/problems/combinations/)（微软、亚马逊、谷歌在半年内面试中考过）
- [全排列](https://leetcode-cn.com/problems/permutations/)（字节跳动在半年内面试常考）
- [全排列 II ](https://leetcode-cn.com/problems/permutations-ii/)（亚马逊、字节跳动、Facebook 在半年内面试中考过）

## 分治 回溯

## 实战题目

- [Pow(x, n) ](https://leetcode-cn.com/problems/powx-n/)（Facebook 在半年内面试常考）
- [子集](https://leetcode-cn.com/problems/subsets/)（Facebook、字节跳动、亚马逊在半年内面试中考过）

## 参考链接

- [牛顿迭代法原理](http://www.matrix67.com/blog/archives/361)
- [牛顿迭代法代码](http://www.voidcn.com/article/p-eudisdmk-zm.html)

## 实战题目

- [多数元素](https://leetcode-cn.com/problems/majority-element/description/) （亚马逊、字节跳动、Facebook 在半年内面试中考过）
- [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)（亚马逊在半年内面试常考）
- [N 皇后](https://leetcode-cn.com/problems/n-queens/)（字节跳动、苹果、谷歌在半年内面试中考过）

