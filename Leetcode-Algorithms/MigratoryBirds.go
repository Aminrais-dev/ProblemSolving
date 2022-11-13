package leetcodealgorithms

import "sort"

/*
Given an array of bird sightings where every element represents a bird type id, determine the id of the most frequently sighted type.
If more than 1 type has been spotted that maximum amount, return the smallest of their ids.
*/

func migratoryBirds(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var count int32
	var sum int32
	var res int32

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			count++
			if count > sum {
				sum = count
				res = arr[i]
			}
		} else {
			count = 0
		}

	}

	return res
}
