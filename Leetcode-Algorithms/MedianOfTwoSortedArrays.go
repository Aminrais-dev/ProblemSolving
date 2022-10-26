package leetcodealgorithms

import "sort"

/* [leetcode]
Difficulty : Hard

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n))
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	median := float64(len(nums1)) / 2

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	if len(nums1)%2 == 1 {
		return float64(nums1[int(median)])
	} else {
		medianA := median + 0.5
		medianB := median - 0.5
		return (float64(nums1[int(medianA)]) + float64(nums1[int(medianB)])) / 2
	}

}
