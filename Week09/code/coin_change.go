package dp

import (
	"math"
)

// 322. 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//
//
// 示例 1:
//
// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3
// 解释: 11 = 5 + 5 + 1
// 示例 2:
//
// 输入: coins = [2], amount = 3
// 输出: -1
//
//
// 说明:
// 你可以认为每种硬币的数量是无限的。

func CoinChange(coins []int, amount int) int {
	var dp func(int) int
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		res := math.MaxInt32
		for _, coin := range coins {
			subSolution := dp(amount - coin)
			if subSolution == -1 {
				continue
			}
			res = min(res, subSolution+1)
		}
		if res != math.MaxInt32 {
			return res
		}
		return -1
	}

	return dp(amount)
}

func CoinChange2(coins []int, amount int) int {
	var dp func(int) int
	memo := make(map[int]int, len(coins))
	dp = func(amount int) int {
		if n, ok := memo[amount]; ok {
			return n
		}
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		res := math.MaxInt32
		for _, coin := range coins {
			subSolution := dp(amount - coin)
			if subSolution == -1 {
				continue
			}
			res = min(res, subSolution+1)
		}
		if res != math.MaxInt32 {
			memo[amount] = res
			return memo[amount]
		}
		return -1
	}

	return dp(amount)
}

func CoinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	// base case
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i < coin {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
