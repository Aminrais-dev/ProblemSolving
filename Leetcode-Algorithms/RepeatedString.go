package leetcodealgorithms

/* [Hackkerrank]
Difficulty : Easy

There is a string 's' of lowercase English letters that is repeated infinitely many times.

Given an integer 'n' find and print the number of letter a's in the first n letters of the infinite string.
*/

func repeatedString(s string, n int64) int64 {
	if s == "a" {
		return n
	}

	var div = n / int64(len(s))
	var modulo = int(n) % (len(s))
	var res int64
	for key, v := range s {
		if string(v) == "a" && key < modulo {
			res += div + 1
		} else if string(v) == "a" {
			res += div
		}
	}

	return res
}
