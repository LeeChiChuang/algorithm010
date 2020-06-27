

# 学习笔记

前序遍历(Pre-order)： 根-左-右

中序遍历(In-order): 左根右

后序遍历(Post-order): 左右根

## **哈希表、映射、集合**

**Hash table（哈希表 散列表） 和 Hash Function（散列函数）**

**完整结构**

![0B6E51B3-A6F2-473D-A02F-AAE634A3A234](https://tva1.sinaimg.cn/large/007S8ZIlly1gg0zw2pi7hj30mb0argnr.jpg)

### Set 不重复的元素



## **树、二叉树、二叉搜索树**

**二叉搜索树，也称为二叉排序树，有序二叉树（Order binary Tree）排序二叉树（Sort Binary Tree），指空树或者具有以下性质的二叉树：**

1. **左子树的所有节点的值均小于它的根节点**
2. **右子树上所有节点的值大于他的根节点**
3. **左右子树也均为二叉查找树**



## 图

### 图的属性

- Grap(V, E)
- V - vertex: 点
  1. 度 - 入度和出度
  2. 点与点之间：连通与否
- E - edge 边
  1. 有向和无向
  2. 权重（边长）

### 图的分类

- 无向无权图
- 有向无权图
- 无向有权图

### 基于图的常见算法

- DFS 递归写法
- BFS 

### 图的高级算法

- [连通图的个数](https://leetcode-cn.com/problems/number-of-islands/)
- 拓扑排序
- 最短路径
- 最小生成树

## Golang哈希Map总结(未完)

### 初始化

Compile/internal/gc/sinit.go

```GO
func maplit(n *Node, m *Node, init *Nodes) {
	// make the map var
	a := nod(OMAKE, nil, nil)
	a.Esc = n.Esc
	a.List.Set2(typenod(n.Type), nodintconst(int64(n.List.Len())))
	litas(m, a, init)

	entries := n.List.Slice()

	// The order pass already removed any dynamic (runtime-computed) entries.
	// All remaining entries are static. Double-check that.
	for _, r := range entries {
		if !isStaticCompositeLiteral(r.Left) || !isStaticCompositeLiteral(r.Right) {
			Fatalf("maplit: entry is not a literal: %v", r)
		}
	}

	if len(entries) > 25 {
		// For a large number of entries, put them in an array and loop.

		// build types [count]Tindex and [count]Tvalue
		tk := types.NewArray(n.Type.Key(), int64(len(entries)))
		te := types.NewArray(n.Type.Elem(), int64(len(entries)))

		tk.SetNoalg(true)
		te.SetNoalg(true)

		dowidth(tk)
		dowidth(te)

		// make and initialize static arrays
		vstatk := staticname(tk)
		vstatk.MarkReadonly()
		vstate := staticname(te)
		vstate.MarkReadonly()

		datak := nod(OARRAYLIT, nil, nil)
		datae := nod(OARRAYLIT, nil, nil)
		for _, r := range entries {
			datak.List.Append(r.Left)
			datae.List.Append(r.Right)
		}
		fixedlit(inInitFunction, initKindStatic, datak, vstatk, init)
		fixedlit(inInitFunction, initKindStatic, datae, vstate, init)

		// loop adding structure elements to map
		// for i = 0; i < len(vstatk); i++ {
		//	map[vstatk[i]] = vstate[i]
		// }
		i := temp(types.Types[TINT])
		rhs := nod(OINDEX, vstate, i)
		rhs.SetBounded(true)

		kidx := nod(OINDEX, vstatk, i)
		kidx.SetBounded(true)
		lhs := nod(OINDEX, m, kidx)

		zero := nod(OAS, i, nodintconst(0))
		cond := nod(OLT, i, nodintconst(tk.NumElem()))
		incr := nod(OAS, i, nod(OADD, i, nodintconst(1)))
		body := nod(OAS, lhs, rhs)

		loop := nod(OFOR, cond, incr)
		loop.Nbody.Set1(body)
		loop.Ninit.Set1(zero)

		loop = typecheck(loop, ctxStmt)
		loop = walkstmt(loop)
		init.Append(loop)
		return
	}
	// For a small number of entries, just add them directly.

	// Build list of var[c] = expr.
	// Use temporaries so that mapassign1 can have addressable key, elem.
	// TODO(josharian): avoid map key temporaries for mapfast_* assignments with literal keys.
	tmpkey := temp(m.Type.Key())
	tmpelem := temp(m.Type.Elem())

	for _, r := range entries {
		index, elem := r.Left, r.Right

		setlineno(index)
		a := nod(OAS, tmpkey, index)
		a = typecheck(a, ctxStmt)
		a = walkstmt(a)
		init.Append(a)

		setlineno(elem)
		a = nod(OAS, tmpelem, elem)
		a = typecheck(a, ctxStmt)
		a = walkstmt(a)
		init.Append(a)

		setlineno(tmpelem)
		a = nod(OAS, nod(OINDEX, m, tmpkey), tmpelem)
		a = typecheck(a, ctxStmt)
		a = walkstmt(a)
		init.Append(a)
	}

	a = nod(OVARKILL, tmpkey, nil)
	a = typecheck(a, ctxStmt)
	init.Append(a)
	a = nod(OVARKILL, tmpelem, nil)
	a = typecheck(a, ctxStmt)
	init.Append(a)
}
```



使用make创建map

```go
// 使用map创建hash golang会在类型检查期间将他们转化成对makemap的调用
// makemap implements Go map creation for make(map[k]v, hint).
// If the compiler has determined that the map or the first bucket
// can be created on the stack, h and/or bucket may be non-nil.
// If h != nil, the map can be created directly in h.
// If h.buckets != nil, bucket pointed to can be used as the first bucket.
func makemap(t *maptype, hint int, h *hmap) *hmap {
	mem, overflow := math.MulUintptr(uintptr(hint), t.bucket.size)
	// 先判断是否溢出或者超出分配的最大值
	if overflow || mem > maxAlloc {
		hint = 0
	}

	// initialize Hmap
	if h == nil {
		h = new(hmap)
	}
	// 获取随机的hash种子
	h.hash0 = fastrand()

  // 根据传入的hint计算最小需要的桶数量
	// Find the size parameter B which will hold the requested # of elements.
	// For hint < 0 overLoadFactor returns false since hint < bucketCnt.
	B := uint8(0)
	for overLoadFactor(hint, B) {
		B++
	}
	h.B = B

	// allocate initial hash table
	// if B == 0, the buckets field is allocated lazily later (in mapassign)
	// If hint is large zeroing this memory could take a while.
  // makeBucketArray 创建用于保存桶的数组
	if h.B != 0 {
		var nextOverflow *bmap
		h.buckets, nextOverflow = makeBucketArray(t, h.B, nil)
		if nextOverflow != nil {
			h.extra = new(mapextra)
			h.extra.nextOverflow = nextOverflow
		}
	}

	return h
}
```

