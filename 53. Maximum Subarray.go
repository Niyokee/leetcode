package main

import "math"

func crossSum(nums []int, left, right, p int) int {
	if left == right {
		return nums[left]
	}

	leftSubSum := math.MinInt32
	currSum := 0

	for i := p; i > left-1; i-- {
		currSum += nums[i]
		leftSubSum = int(math.Max(float64(leftSubSum), float64(currSum)))
	}

	rightSubSum := math.MinInt32
	currSum = 0
	for i := p; i < right+1; i++ {
		currSum += nums[i]
		rightSubSum = int(math.Max(float64(rightSubSum), float64(currSum)))
	}

	return rightSubSum + leftSubSum
}

func helper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	p := (left + right) / 2
	leftSum := float64(helper(nums, left, p))
	rightSum := float64(helper(nums, p+1, right))
	crossSum := float64(crossSum(nums, left, right, p))

	return int(math.Max(math.Max(leftSum, rightSum), crossSum))

}

func maxSubArray(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}
