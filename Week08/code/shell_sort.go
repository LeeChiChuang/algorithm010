package sort

// 将序列分割成若干子序列进行直接插入排序
func ShellSort(nums []int) []int {
	numsLen := len(nums)
	for gap := numsLen / 2; gap > 0; gap = gap / 2 {
		for i:=gap; i<numsLen; i++ {
			j := i
			current := nums[i]
			for j-gap >= 0 && current < nums[j-gap] {
				nums[j] = nums[j-gap]
				j = j-gap
			}
			nums[j] = current
		}
	}

	return nums
}
