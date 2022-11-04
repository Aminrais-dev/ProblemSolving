package leetcodealgorithms

import "math"

/*
Watson likes to challenge Sherlock's math ability. He will provide a starting and ending value that describe a range of integers,
inclusive of the endpoints. Sherlock must determine the number of square integers within that range.
*/

func squares(a int32, b int32) int32 {

	var count int32
	var max = int32(math.MaxInt32)
	for i := int32(1); i <= max; i++ {
		var square = i * i
		if square < a {
			continue
		}

		if square >= a && square <= b {
			count++
		} else if square > b {
			break
		}

	}

	return count

}
