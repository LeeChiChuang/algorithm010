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







