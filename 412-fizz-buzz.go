package main

import "strconv"

func fizzBuzz(n int) []string {
	var ans []string
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			ans = append(ans, "FizzBuzz")
		case i%3 == 0:
			ans = append(ans, "Fizz")
		case i%5 == 0:
			ans = append(ans, "Buzz")
		default:
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}
