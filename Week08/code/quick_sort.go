package sort

import (
	"fmt"
	"math/rand"
)

func QuickSort(nums []int, begin, end int) []int {
	if end <= begin {
		return nil
	}
	pivot := partition(nums, begin, end)
	QuickSort(nums, begin, pivot-1)
	QuickSort(nums, pivot+1, end)

	return nums
}

// 7, 8, 9, 2, 3, 5
func partition(nums []int, begin, end int) int {
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if nums[i] < nums[pivot] {
			temp := nums[counter]
			nums[counter] = nums[i]
			nums[i] = temp
			counter++
		}
	}
	fmt.Println(nums)

	temp := nums[pivot]
	nums[pivot] = nums[counter]
	nums[counter] = temp
	fmt.Println(nums)

	return counter
}

func QuickSort2(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort2(a[:left])
	QuickSort2(a[left+1:])

	return a
}