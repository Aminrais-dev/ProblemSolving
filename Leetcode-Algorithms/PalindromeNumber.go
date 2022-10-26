package leetcodealgorithms

import "strconv"

/* [leetcode]
Difficulty : Easy

Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.
*/

func isPalindrome(x int) bool {
	var cek = strconv.Itoa(x)
	var palindrome string
	for i := len(cek) - 1; i >= 0; i-- {
		palindrome += string(cek[i])
	}

	return cek == palindrome
}
