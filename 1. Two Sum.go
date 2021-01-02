package main

/*
My first brute force answer
Time: O(n^2)
Space: O(1)
*/
func twoSum(nums []int, target int) []int {
	for i, n1 := range nums {
		j := i + 1
		for j < len(nums) {
			n2 := nums[j]
			if n1+n2 == target {
				return []int{i, j}
			}
			j++
		}
	}
	return []int{}
}
