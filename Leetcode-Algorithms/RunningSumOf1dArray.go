package leetcodealgorithms

/* [leetcode]
Difficulty : Easy

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.
*/

func runningSum(nums []int) []int {
	var res []int
	var nilai int
	for _, v := range nums {
		nilai += v
		res = append(res, nilai)
	}

	return res
}
