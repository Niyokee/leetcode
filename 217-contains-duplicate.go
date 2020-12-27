package main

func containsDuplicate(nums []int) bool {
	h := make(map[int]bool)
	for _, n := range nums {
		if _, ok := h[n]; ok {
			return true
		}
		h[n] = true
	}
	return false
}
