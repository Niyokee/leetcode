package main

/*
38を超えるとacceptされなくなる[Wrong Answer]
*/
func climbStairs_wrong(n int) int {
	// 2の数を求める
	var ans int = 1
	for two := n / 2; two > 0; two-- {
		ones := n - (two * 2)
		ans += combination(two+ones, two)
	}
	return ans
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

/*
Approach 1: Brute Force[Time Limit exceed]
Time Complexity: 2^n
Space Complexity:O(n)
*/

func climbStairs1(n int) int {
	return climb_Stairs1(0, n)
}

func climb_Stairs1(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return climb_Stairs1(i+1, n) + climb_Stairs1(i+2, n)
}

/*
Approach3 Dynamic Programing
Time Complexity: O(n)
Space Complexity:O(n)
*/

func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
