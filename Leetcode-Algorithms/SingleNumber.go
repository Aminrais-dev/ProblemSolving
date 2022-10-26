package leetcodealgorithms

/*
[leetcode]

Difficulty : Easy

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.
*/
func singleNumber(nums []int) int {
	maps := map[int]int{}
	for _, v := range nums {
		maps[v]++
	}
	for key, v := range maps {
		if v == 1 {
			return key
		}
	}

	return 0
}
