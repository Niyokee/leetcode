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

/*
上の回答の２つのif文がforの中に入っている必要はないので外に出した。
*/
func missingNumber1(nums []int) int {
	sort.Ints(nums)
	// ensure that 0 is the first index
	if nums[0] != 0 {
		return 0
	} else if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}
	for i, n := range nums {
		if n+1 != nums[i+1] {
			return n + 1
		}
	}
	return 0
}
