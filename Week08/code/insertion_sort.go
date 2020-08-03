package sort

func InsertionSort(nums []int) []int {
	length := len(nums)
	for i:=1; i<length; i++ {
		prevIndex := i-1
		curr := nums[i]
		for prevIndex>=0 && nums[prevIndex] > curr {
			nums[prevIndex+1] = nums[prevIndex]
			prevIndex--
		}
		nums[prevIndex+1] = curr
	}

	return nums
}
