package sort

func SelectionSort(nums []int) []int {
	length := len(nums)
	for i:=0; i<length-1; i++ {
		minIndex := i
		for j:=i+1; j<length; j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		tmp := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = tmp
	}

	return nums
}
