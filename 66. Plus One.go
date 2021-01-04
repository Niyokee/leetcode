package main

func plusOne(digits []int) []int {
	for l := len(digits) - 1; l >= 0; l-- {
		if digits[l] < 9 {
			digits[l]++
			return digits
		} else if l == 0 {
			digits[l] = 0
			digits = append([]int{1}, digits...)
			return digits
		}
		digits[l] = 0
	}
	return digits
}
