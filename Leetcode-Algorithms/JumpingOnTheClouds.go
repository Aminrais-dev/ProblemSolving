package leetcodealgorithms

/*
here is a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. The player can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus 1 or 2.
The player must avoid the thunderheads. Determine the minimum number of jumps it will take to jump from the starting postion to the last cloud. It is always possible to win the game.

For each game, you will get an array of clouds numbered 0 if they are safe or 1 if they must be avoided.
*/

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
