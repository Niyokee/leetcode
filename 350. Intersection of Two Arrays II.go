package main

import "math"

func intersect(nums1 []int, nums2 []int) []int {
	h1 := make(map[int]int)
	h2 := make(map[int]int)
	var ans []int

	for _, n := range nums1 {
		h1[n]++
	}
	for _, n := range nums2 {
		h2[n]++
	}

	for k, v := range h1 {
		if _, ok := h2[k]; !ok {
			continue
		}
		if v == h2[k] {
			for i := 0; i < v; i++ {
				ans = append(ans, k)
			}
		} else {
			n := int(math.Min(float64(v), float64(h2[k])))
			for n > 0 {
				ans = append(ans, k)
				n--
			}
		}
	}
	return ans
}
