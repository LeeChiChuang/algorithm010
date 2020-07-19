package dp

import "math"

// https://leetcode-cn.com/problems/maximum-subarray/
// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 示例:
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 进阶:
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

func MaxSumArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[1]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := math.MinInt32
	for i:=1; i<len(nums); i++ {
		dp[i] = max(nums[i], nums[i] + dp[i-1])
		res = max(res, dp[i])
	}

	return res
}

// 状态压缩 降低空间复杂度

func MaxSumArray2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := nums[0]
	prev := dp
	res := prev
	for i:=1; i<len(nums); i++ {
		dp = max(nums[i], nums[i] + prev)
		res = max(res, dp)
		prev = dp
	}

	return res
}