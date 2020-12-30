package main

import "fmt"

func hammingWeight(num uint32) int {
	var bits int
	var mask uint32 = 1

	for i := 0; i < 32; i++ {
		fmt.Println(i, num, mask)
		if num&mask != 0 {
			bits++
		}
		mask <<= 1
	}

	return bits
}
