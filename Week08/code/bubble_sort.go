package sort

func BubbleSort(nums []int) []int {
	n := len(nums)
	for i:=0; i<n - 1; i++ {
		for j:=0; j<n-1-i; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}

	return  nums
}