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

/*
find second index more efficiently to put key and value to hash map
Time: O(n)
Space: O(n)
*/
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, j := range nums {
		m[j] = i
	}
	for i, _ := range nums {
		t := target - nums[i]
		if v, ok := m[t]; ok && v != i {
			return []int{i, v}
		}
	}
	return []int{}
}

func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	for i, j := range nums {
		t := target - nums[i]
		if v, ok := m[t]; ok && v != i {
			return []int{i, v}
		}
		m[j] = i
	}
	return []int{}
}
