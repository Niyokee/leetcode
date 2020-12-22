package main

/*
two pointer approach
time:  O(N) to swap N/2 element
space: O(1)
*/
func reverseString(s []byte) {
	var firstPointer int
	secondPointer := len(s) - 1

	for firstPointer < secondPointer {
		s[secondPointer], s[firstPointer] = s[firstPointer], s[secondPointer]
		firstPointer++
		secondPointer--
	}
}
