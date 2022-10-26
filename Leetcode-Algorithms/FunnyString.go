package leetcodealgorithms

import "math"

/* [Hackkerrank]
Difficulty : Easy

In this challenge, you will determine whether a string is funny or not. To determine whether a string is funny, create a copy of the string in reverse e.g. .
Iterating through each string, compare the absolute difference in the ascii values of the characters at positions 0 and 1, 1 and 2 and so on to the end.
If the list of absolute differences is the same for both strings, they are funny.

Determine whether a give string is funny. If it is, return Funny, otherwise return Not Funny.
*/

func funnyString(s string) string {

	var cek []int
	for i := 0; i < len(s)-1; i++ {
		var abs = int(math.Abs(float64(s[i]) - float64(s[i+1])))
		cek = append(cek, abs)
	}

	var index int
	for i := len(s) - 1; i > 0; i-- {
		var abs = int(math.Abs(float64(s[i]) - float64(s[i-1])))
		if abs != cek[index] {
			return "Not Funny"
		}
		index++
	}

	return "Funny"

}
