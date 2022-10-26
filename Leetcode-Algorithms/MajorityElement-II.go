package leetcodealgorithms

/* [leetcode]
Difficulty : Medium

Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
*/

func majorityElement(nums []int) []int {
	maps := map[int]int{}
	for _, v := range nums {
		maps[v]++
	}

	var res []int
	for key, val := range maps {
		if val > len(nums)/3 {
			res = append(res, key)
		}
	}

	return res
}
