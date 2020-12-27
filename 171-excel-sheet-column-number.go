package main

import (
	"math"
)

func titleToNumber(s string) int {
	var ans int
	l := len(s) - 1
	for i, r := range s {
		// get ascii code
		ascii := int(r) - 64
		// get exponential
		n := l - i
		if n == 0 {
			ans += ascii
		} else {
			ans += int(math.Pow(float64(26), float64(n))) * ascii
		}
	}
	return ans
}
