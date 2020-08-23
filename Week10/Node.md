# DFS & BSF



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

