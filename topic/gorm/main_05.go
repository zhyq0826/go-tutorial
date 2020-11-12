package main

import "fmt"

func bsearch(nums []int, target int) int {
	low := 0
	hi := len(nums) - 1
	for low <= hi {
		mid := low + (hi-low)/2
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		}
	}

	return -1
}

func main() {
	// nums := []int{1, 3, 4, 6, 8, 9, 10, 11, 12, 12, 15, 20}
	nums := []int{1}
	fmt.Println(bsearch(nums, 1))
}
