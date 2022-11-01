package leetcodealgorithms

/*Hackerrank [Easy]
Reduce a string of lowercase characters in range ascii[‘a’..’z’]by doing a series of operations. In each operation,
select a pair of adjacent letters that match, and delete them.

Delete as many characters as possible using this method and return the resulting string. If the final string is empty,
return Empty String
*/

func superReducedString(s string) string {

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = string(s[:i]) + string(s[i+2:])
			i = -1
		}
	}

	if s == "" {
		return "Empty String"
	}

	return s

}
