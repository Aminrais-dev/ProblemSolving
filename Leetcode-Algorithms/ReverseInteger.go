package leetcodealgorithms

import (
	"math"
	"strconv"
)

/*
Difficulty : Medium

Given a signed 32-bit integer x, return x with its digits reversed.

If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
*/

func reverse(x int) int {

	var temp string
	if x < 0 {
		temp += "-"
	}

	var abs = int(math.Abs(float64(x)))
	for abs != 0 {
		rever := abs % 10
		temp += strconv.Itoa(rever)
		abs /= 10
	}

	res, _ := strconv.Atoi(temp)
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res

}
