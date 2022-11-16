package leetcodealgorithms

func jumpingOnClouds(c []int32) int32 {

	var res int32
	for i := 0; i < len(c); i++ {
		if i+2 < len(c)-1 {
			if c[i+2] != 1 {
				res++
				i = i + 1
			} else if c[i+1] != 1 {
				res++
			}
		}

	}

	return res + 1

}
