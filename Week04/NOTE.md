# 深度优先搜索和广度优先搜索



## DFS代码模板

> 递归写法

```python
visited = set()

def dfs(node, visited):
  if node in visited: # terminator
    return
  
  visited.add(node)
  
  for next_node in node.childre():
    if not next_node in visited:
      dfs(next_node, visited)
```

> 非递归写法

```python
def DFS(self, tree):
  if tree.root is Node:
    return []
  
  visited, stack = [], [tree.root]
  while stack:
    node = stack.pop()
    visited.add(node)
    
    process(node)
    nodes = generate_ralated_nodes(node)
    stack.pus(nodes)
```



## 广度优先搜索

> 非递归

```python
def BFS(graph, start, end):
  
  queue = []
  queue.append([start])
  visited.add(start)
  
  while queue:
    node = queue.pop
    visited.add(node)
    
    process(node)
    node = generate_ralated_nodes(node)
    queue.push(nodes)
```



# 贪心算法

## 贪心算法Greedy

>贪心算法是一种在每一步选择中都采取在当前状态下最好或者最优(即最有利)的选择，从而希望导致结果是全局最好或最优的算法。
>
>贪心和动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能。
>
>贪心可以解决一些最优问题，如：求图中的最小生成树、求哈弗曼编码等。然而对于工程和生活中的问题，贪心一般不能得到我们想要的答案。
>
>一旦问题可以通过贪心来解决，那么贪心法一般是解决这个问题的最好办法。由于贪心的高效性以及其求得的答案比较接近最优结果，贪心法也可以用作辅助算法或者直接解决一些要求不特别精确的问题。



## 贪心法的场景

简单地说，问题能过分解成子问题来解决，子问题的最优解能递归推到最终问题的最优解。这个问题最优解称为最优子结构。



# 二分查找

## 前提

1. 目标函数单调性
2. 存在上下界 bounde
3. 能够通过索引访问

## 代码模板

```python
left, right = 0, len(array) -1
while lef <= right:
  mid = (left <= right)
  if array[mid] == target:
    break or return result
  elif array[mid] < target:
    left = mid+ 1
  else:
    right = mid - 1 
```

## 深度优先搜索和广度优先搜索

## 实战题目

- [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/#/description)（字节跳动、亚马逊、微软在半年内面试中考过）
- [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/#/description)
- [括号生成](https://leetcode-cn.com/problems/generate-parentheses/#/description)（字节跳动、亚马逊、Facebook 在半年内面试中考过）
- [在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/#/description)（微软、亚马逊、Facebook 在半年内面试中考过）

## 课后作业

- [单词接龙](https://leetcode-cn.com/problems/word-ladder/description/)（亚马逊在半年内面试常考）
- [单词接龙 II ](https://leetcode-cn.com/problems/word-ladder-ii/description/)（微软、亚马逊、Facebook 在半年内面试中考过）
- [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)（近半年内，亚马逊在面试中考查此题达到 350 次）
- [扫雷游戏](https://leetcode-cn.com/problems/minesweeper/description/)（亚马逊、Facebook 在半年内面试中考过）

## 贪心算法

## 参考链接

- [coin change 题目](https://leetcode-cn.com/problems/coin-change/)
- [动态规划定义](https://zh.wikipedia.org/wiki/动态规划)

## 课后作业

- [柠檬水找零](https://leetcode-cn.com/problems/lemonade-change/description/)（亚马逊在半年内面试中考过）
- [买卖股票的最佳时机 II ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/)（亚马逊、字节跳动、微软在半年内面试中考过）
- [分发饼干](https://leetcode-cn.com/problems/assign-cookies/description/)（亚马逊在半年内面试中考过）
- [模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/description/)
- [跳跃游戏](https://leetcode-cn.com/problems/jump-game/) （亚马逊、华为、Facebook 在半年内面试中考过）
- [跳跃游戏 II ](https://leetcode-cn.com/problems/jump-game-ii/)（亚马逊、华为、字节跳动在半年内面试中考过）

## 二分查找

## 参考链接

- [二分查找代码模板](https://shimo.im/docs/xvIIfeEzWYEUdBPD/)
- [Fast InvSqrt() 扩展阅读](https://www.beyond3d.com/content/articles/8/)

## 实战题目

- [x 的平方根](https://leetcode-cn.com/problems/sqrtx/)（字节跳动、微软、亚马逊在半年内面试中考过）
- [有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)（亚马逊在半年内面试中考过）

## 课后作业

- [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)（Facebook、字节跳动、亚马逊在半年内面试常考）
- [搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)（亚马逊、微软、Facebook 在半年内面试中考过）
- [寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)（亚马逊、微软、字节跳动在半年内面试中考过）
- 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方
  说明：同学们可以将自己的思路、代码写在学习总结中