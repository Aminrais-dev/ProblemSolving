package leetcodealgorithms

import "math"

/*
The distance between two array values is the number of indices between them.
Given a, find the minimum distance between any pair of equal elements in the array. If no such value exists, return -1 .
*/

func minimumDistances(a []int32) int32 {

	min := math.MaxInt32
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				temp := j - i
				if temp < min {
					min = temp
				}
			}
		}
	}
	if min == math.MaxInt32 {
		min = -1
	}

	return int32(min)
}
