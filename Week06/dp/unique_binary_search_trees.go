package dp

// 96. 不同的二叉搜索树
// 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
// 示例:
//
// 输入: 3
// 输出: 5
// 解释:
// 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
// 1         3     3      2      1
// \       /     /      / \      \
// 3     2     1      1   3      2
// /     /       \                 \
// 2     1         2                 3

func NumTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}

	return G[n]
}

// 数学公式 卡塔兰数
func NumTrees2(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2 * i + 1) / (i + 2)
	}

	return c
}