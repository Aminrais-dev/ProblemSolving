package leetcodealgorithms

import (
	"fmt"
	"sort"
)

/*
There is a large pile of socks that must be paired by color.
Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.
*/

func sockMerchant(n int32, ar []int32) int32 {

	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	fmt.Println(ar)
	var count = 1
	var res int32
	for i := 0; int32(i) < n-1; i++ {
		fmt.Println(ar[i], ar[i+1])
		if ar[i] == ar[i+1] {
			count++
			fmt.Println(count)
			if count > 1 && count%2 == 0 {
				res++
				fmt.Println(res)
			}
		} else {
			count = 1
		}

	}

	return res

}
