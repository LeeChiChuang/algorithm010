package dp

import (
	"math"
)

// https://leetcode-cn.com/problems/minimum-path-sum/
// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例:
//
// 输入:
// [
// [1,3,1],
// [1,5,1],
// [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。

func MinPathSum(grid [][]int) int {
	return pathDfs(grid, 0, 0)
}

func pathDfs(grid [][]int, i, j int) int {
	if i == len(grid) || j == len(grid[0]) {
		return math.MaxInt32
	}
	if i == len(grid)-1 && j == len(grid[0])-1 {
		return grid[i][j]
	}

	right := pathDfs(grid, i, j+1)
	down := pathDfs(grid, i+1, j)
	return min(right, down) + grid[i][j]
}

// 二维动态规划

func MinPathSum2(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i:=len(grid)-1; i>=0; i-- {
		dp[i] = make([]int, len(grid[i]))
		for j:=len(grid[0]) - 1; j>=0; j-- {
			if i == len(grid) - 1 && j != len(grid[0]) - 1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else if j == len(grid[0]) - 1 && i != len(grid) - 1 {
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else if i != len(grid) - 1 && j != len(grid[0]) - 1 {
				dp[i][j] = grid[i][j] + min(dp[i][j+1], dp[i+1][j])
			} else {
				dp[i][j] = grid[i][j]
			}
		}
	}

	return  dp[0][0]
}