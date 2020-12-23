package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i+2 < len(nums); i += 2 {
		fmt.Println(nums[i], nums[i+1])
		if nums[i] != nums[i+1] {
			return nums[i]
		}

	}
	return nums[len(nums)-1]
}
