package main

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, n := range nums {
		if i == 0 && n != 0 {
			return 0
		}
		if i+1 == len(nums) {
			return nums[len(nums)-1] + 1
		}
		if n+1 != nums[i+1] {
			return n + 1
		}
	}
	return 0
}
