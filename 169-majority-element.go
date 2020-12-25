package main

import (
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//func majorityElement(nums []int) int {
//	h := make(map[int]int)
//
//	border := math.Ceil(float64(len(nums)) / 2.0)
//
//	for _, n := range nums {
//		if _, ok := h[n]; ok {
//			h[n]++
//			if float64(h[n]) >= border { return n }
//		} else {
//			h[n] = 1
//			if float64(h[n]) >= border { return n }
//		}
//
//	}
//	return 0
//}
