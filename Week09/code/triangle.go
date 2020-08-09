package dp

// 120. 三角形最小路径和
// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//
//
// 例如，给定三角形：
//
// [
// [2],
// [3,4],
// [6,5,7],
// [4,1,8,3]
// ]
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//
//
// 说明：
//
// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

func MinimumTotal(triangle [][]int) int {
	return dfs(triangle, 0, 0)
}

// i 行 j 列
func dfs(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}
	// log.Printf("i:%d \n", i)
	// log.Printf("j:%d \n", j)
	dfs1 := dfs(triangle, i+1, j)
	dfs2 := dfs(triangle, i+1, j+1)
	// log.Printf("dfs1:%d \n", dfs1)
	// log.Printf("dfs2:%d \n", dfs2)
	ret := min(dfs1, dfs2)
	sum := triangle[i][j]
	// log.Printf("ret i:%d \n", i)
	// log.Printf("ret j:%d \n", j)
	// log.Printf("ret:%d \n", ret)
	// log.Printf("sum:%d \n", sum)
	return ret + sum
	/*
		for ; i < len(triangle); i++ {
			tmp := i + 1
			for ; i < tmp; j++  {
				val = triangle[i] + triangle[i][j]
				recursion(triangle, i, j, val)
			}
		}
	*/
}
