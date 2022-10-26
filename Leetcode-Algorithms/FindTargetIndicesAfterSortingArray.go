package leetcodealgorithms

import "sort"

/* [leetcode]
Difficulty : Easy

You are given a 0-indexed integer array nums and a target element target.

A target index is an index i such that nums[i] == target.

Return a list of the target indices of nums after sorting nums in non-decreasing order.
If there are no target indices, return an empty list. The returned list must be sorted in increasing order.
*/

func targetIndices(nums []int, target int) []int {

	sort.Ints(nums)

	var res []int
	for key, v := range nums {
		if v > target {
			break
		} else if v < target {
			continue
		}
		res = append(res, key)
	}

	return res

}
