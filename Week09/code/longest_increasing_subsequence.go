package dp

// 300. 最长上升子序列
// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
// 示例:
//
// 输入: [10,9,2,5,3,7,101,18]
// 输出: 4
// 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
// 说明:
//
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
// 你算法的时间复杂度应该为 O(n2) 。
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

// 数学归纳法 -> 状态转移方程
// 1. 明确dp数组所存数据的含义
// 2. 运用数学归纳法的思想，假设dp[0...i-1]都已知，想办法求出dp[i]

func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for idx, _:= range dp  {
		dp[idx] = 1
	}

	for i:=0; i<len(nums); i++ {
		for j:=0; j<i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	res := 0
	for _, val := range dp {
		res = max(res, val)
	}

	return res
}

// 二分查找法 patience game
func LengthOfLIS2(nums []int) int {
	return 0
}
