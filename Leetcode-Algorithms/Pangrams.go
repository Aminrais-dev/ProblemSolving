package leetcodealgorithms

import (
	"fmt"
	"strings"
)

/*
A pangram is a string that contains every letter of the alphabet. Given a sentence determine whether it is a pangram in the English alphabet.
Ignore case. Return either pangram or not pangram as appropriate.
*/
func pangrams(s string) string {

	var count int
	var cek = "abcdefghijklmnopqrstuvwxyz"
	var cek2 = strings.ToLower(s)
	if len(s) < 26 {
		return "not pangram"
	}
	for i := 0; i < len(cek); i++ {
		for j := 0; j < len(cek2); j++ {
			if cek[i] == cek2[j] {
				count++
				break
			}
		}
	}

	fmt.Println(count, len(cek))
	if count == len(cek) {
		return "pangram"
	} else {

		return "not pangram"
	}

}
