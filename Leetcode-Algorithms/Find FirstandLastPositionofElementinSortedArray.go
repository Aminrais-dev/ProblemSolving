package leetcodealgorithms

func searchRange(nums []int, target int) []int {

	var res []int
	for key, v := range nums {
		if v > target {
			break
		} else if v < target {
			continue
		}
		res = append(res, key)
	}

	if res == nil {
		return []int{-1, -1}
	} else if len(res) == 1 {
		return append(res, res...)
	}

	last := res[len(res)-1]
	return append(res[:1], last)

}
