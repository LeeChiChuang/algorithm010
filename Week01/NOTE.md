# 学习笔记



## **简单：**

1. 用 add first 或 add last 这套新的 API 改写 Deque 的代码

2. 分析 Queue 和 Priority Queue 的源码

3. [删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)（Facebook、字节跳动、微软在半年内面试中考过）

4. [旋转数组](https://leetcode-cn.com/problems/rotate-array/)（微软、亚马逊、PayPal 在半年内面试中考过）

5. [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)（亚马逊、字节跳动在半年内面试常考）

6. [合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)（Facebook 在半年内面试常考）

**Go语言实现**

> #### 方法一 : 合并后排序

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1 = append(nums1[:m], nums2[:n]...)
    sort.Ints(nums1)
}
```



> #### 方法二：双指针 / 从前往后

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1Copy := make([]int, m)
    copy(nums1Copy, nums1[:m])
    nums1 = nums1[0:0]
    p1 := 0
    p2 := 0

    for p1<m && p2<n {
        if (nums1Copy[p1] < nums2[p2]) {
            nums1 = append(nums1, nums1Copy[p1])
            p1++
        } else {
            nums1 = append(nums1, nums2[p2])
            p2++
        }
    }

    if p1<m {
        nums1 = append(nums1, nums1Copy[p1:]...)
    } else if p2 < n {
            nums1 = append(nums1, nums2[p2:]...)
    }

}
```

> #### 双指针 / 从后往前



7. [两数之和](https://leetcode-cn.com/problems/two-sum/)（亚马逊、字节跳动、谷歌、Facebook、苹果、微软在半年内面试中高频常考）

**一遍哈希表**

```go
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
```



8. [移动零](https://leetcode-cn.com/problems/move-zeroes/)（Facebook、亚马逊、苹果在半年内面试中考过）

   ```python
   class Solution(object):
       def moveZeroes(self, nums):
           """
           :type nums: List[int]
           :rtype: None Do not return anything, modify nums in-place instead.
           """
           j = 0
           for i in range(len(nums)):
               if nums[i] != 0:
                   if i != j:
                       nums[j] = nums[i]
                       nums[i] = 0
                   j = j+1
   
           return nums
   ```

   

9. [加一](https://leetcode-cn.com/problems/plus-one/)（谷歌、字节跳动、Facebook 在半年内面试中考过）



## **中等：**

- [设计循环双端队列](https://leetcode.com/problems/design-circular-deque)（Facebook 在 1 年内面试中考过）

## **困难：**

- [接雨水](https://leetcode.com/problems/trapping-rain-water/)（亚马逊、字节跳动、高盛集团、Facebook 在半年内面试常考）

## 



## LRU缓存机制

[leetcode](https://leetcode-cn.com/problems/lru-cache/)   实现方法 哈希表 + 双向链表  

通过哈希表获取节点所在位置 存在的话将节点移到链表头部

put时先验证key是否存在，存在直接将对应的node修改值移到头部，不存在的话，新建节点，放在头部

判断是否超过容量，超过的话删除链表尾部节点(tail.prev)

**Go语言版本**

```go
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	size int
	head *Node
	tail *Node
	data map[int]*Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size: capacity,
		head: &Node{},
		tail: &Node{},
		data: make(map[int]*Node),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lru *LRUCache) Get(key int) int {
	// 值存在的话 将节点移到头部 返回值
	v, exist := lru.data[key]
	if exist {
		if v != lru.head.next {
			v.next.prev = v.prev
			v.prev.next = v.next

			v.prev = lru.head
			v.next = lru.head.next

			lru.head.next.prev = v
			lru.head.next = v
		}
		return v.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	// 先验证是否存在 存在的话会将当前移到头部 表示最近使用过
	if lru.Get(key) != -1 {
		lru.data[key].val = value
		return
	}

	lru.data[key] = &Node{
		key: key,
		val:  value,
		prev: lru.head,
		next: lru.head.next,
	}
	lru.head.next.prev = lru.data[key]
	lru.head.next = lru.data[key]

	if lru.Len() > lru.size {
		delete(lru.data, lru.tail.prev.key)
		lru.tail.prev = lru.tail.prev.prev
		lru.tail.prev.next = lru.tail
	}
}

func (lru *LRUCache) Len() int {
	return len(lru.data)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```

**Python**

```python
class LRUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.capacity = capacity
        self.hashMap = {}
        self.head = ListNode()
        self.tail = ListNode()
        self.head.next = self.tail
        self.tail.prev = self.head

    def move_node_to_tail(self, key):
        node = self.hashMap[key]
        # 节点前面的执行下一个
        node.prev.next = node.next
        node.next.prev = node.prev
        # 节点移到tail
        node.prev = self.tail.prev
        node.next = self.tail
        self.tail.prev.next = node
        self.tail.prev = node

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        
        if key in self.hashMap:
            self.move_node_to_tail(key)

        res = self.hashMap.get(key, -1)
        if res == -1:
            return res

        return res.value

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """
        """
        if key in self.hashMap:
            self.hashMap[key].value = value
            self.move_node_to_tail(key)
        """
        if self.get(key) != -1 :
            self.hashMap[key].value = value
        else:
            if self.capacity == len(self.hashMap):
                # 去掉对应最近未被使用hashMap对应的值
                self.hashMap.pop(self.head.next.key)
                self.head.next = self.head.next.next
                self.head.next.prev = self.head

            new = ListNode(key, value)
            self.hashMap[key] = new
            new.prev = self.tail.prev
            new.next = self.tail
            self.tail.prev.next = new
            self.tail.prev = new


class ListNode(object):

    def __init__(self, key=None, value=None):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None
```

