package main

import "fmt"

func findFirstPositive(nums []int) int {
	n := len(nums)
	for index := 0; index < n; index++ {
		if nums[index] > 0 && nums[index] <= n && nums[nums[index]-1] != nums[index] {
			nums[nums[index]-1], nums[index] = nums[index], nums[nums[index]-1]
		}
	}

	for index := 0; index < n; index++ {
		if nums[index] != index+1 {
			return index + 1
		}
	}

	return n + 1
}

func main() {
	nums := []int{1, -1, 3, 4}
	fmt.Println(findFirstPositive(nums))
	nums = []int{7, 8, 9, 11, 12}
	fmt.Println(findFirstPositive(nums))
	nums = []int{0, 1, 2, 5, 6, 7, 8}
	fmt.Println(findFirstPositive(nums))
	nums = []int{9, 10, 88, 7, 56, 11, 11}
	fmt.Println(findFirstPositive(nums))
	nums = []int{1, 1, 3, 4}
	fmt.Println(findFirstPositive(nums))
	nums = []int{}
	fmt.Println(findFirstPositive(nums))
}
