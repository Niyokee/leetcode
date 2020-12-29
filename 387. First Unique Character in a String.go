package main

import "strings"

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	h := make(map[string]int)

	// count appearance
	for _, r := range s {
		h[string(r)]++
	}

	var (
		min   int
		count int
	)

	var hasUnique bool
	for k, v := range h {
		// continue it v is not unique
		if v != 1 {
			continue
		} else {
			hasUnique = true
		}

		i := strings.Index(s, k)

		if min >= i || count == 0 {
			min = i
		}
		count++
	}
	if !hasUnique {
		return -1
	} else {
		return min
	}
}

/*
I assumed that map iteration in go did not support its order, so
I compare the index of a value every time.
However if i use string, the parameter I dont have to do the complicated things
*/
func firstUniqChar2(s string) int {
	c := make(map[rune]int)

	for _, r := range s {
		c[r]++
	}

	for i, v := range s {
		if c[v] == 1 {
			return i
		}
	}

	return -1
}
