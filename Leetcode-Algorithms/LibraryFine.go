package leetcodealgorithms

/*
Your local library needs your help! Given the expected and actual return dates for a library book, create a program that calculates the fine (if any). The fee structure is as follows:

If the book is returned on or before the expected return date, no fine will be charged (i.e.: fine = 0).
If the book is returned after the expected return day but still within the same calendar month and year as the expected return date, fine = 15 Hackos * (the number of days late).
If the book is returned after the expected return month but still within the same calendar year as the expected return date, the fine = 500 Hackos * (the number of days late).
If the book is returned after the calendar year in which it was expected, there is a fixed fine of 10000 Hackos.
Charges are based only on the least precise measure of lateness. For example, whether a book is due January 1, 2017 or December 31, 2017, if it is returned January 1, 2018, that is a year late and the fine would be 10000 Hackos.
*/
func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {

	var res int32
	if y1 == y2 {
		if m1 == m2 {
			if d1 > d2 {
				res = (d1 - d2) * 15
			}
		} else if m1 > m2 {
			res = (m1 - m2) * 500
		}
	} else if y1 > y2 {
		res = 10000
	}

	return res

}
