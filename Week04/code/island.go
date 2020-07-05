package main

import "fmt"

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//  
//
// 示例 1:
//
// 输入:
// 11110
// 11010
// 11000
// 00000
// 输出: 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-islands
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	arr := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '1'}}

	ret := numIslands(arr)
	fmt.Println(ret)
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	column := len(grid[0])
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, grid)
				count++
			}
		}
	}

	return count
}

func dfs(i int, j int, grid [][]byte) {
	grid[i][j] = '0'
	if (i+1) < len(grid) && grid[i+1][j] == '1' {
		dfs(i+1, j, grid)
	}
	if (i-1) >= 0 && grid[i-1][j] == '1' {
		dfs(i-1, j, grid)
	}
	if (j+1) < len(grid[i]) && grid[i][j+1] == '1' {
		dfs(i, j+1, grid)
	}
	if (j-1) >= 0 && grid[i][j-1] == '1' {
		dfs(i, j-1, grid)
	}
}
