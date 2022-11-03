package leetcodealgorithms

/*
The Utopian Tree goes through 2 cycles of growth every year. Each spring, it doubles in height. Each summer, its height increases by 1 meter.

A Utopian Tree sapling with a height of 1 meter is planted at the onset of spring. How tall will the tree be after  growth cycles?

For example, if the number of growth cycles is , the calculations are as follows:
*/

func utopianTree(n int32) int32 {

	h := 1
	for i := 0; int32(i) < n; i++ {
		if i%2 == 0 {
			h *= 2
		} else {
			h++
		}
	}

	return int32(h)
}
